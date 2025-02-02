package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

// RequestLogger logs detailed request and response information
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("ENABLE_REQUEST_LOGGING") != "true" {
			next.ServeHTTP(w, r)
			return
		}

		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		log.Printf("Headers: %v", r.Header)

		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil {
				log.Printf("Request Body: %s", string(bodyBytes))
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		rw := &responseWriter{
			ResponseWriter: w,
			body:          new(bytes.Buffer),
		}

		next.ServeHTTP(rw, r)

		log.Printf("Response Status: %d", rw.status)
		log.Printf("Response Headers: %v", w.Header())
		log.Printf("Response Body: %s", rw.body.String())
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
	body   *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
} 