package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// requestLoggerMiddleware logs detailed request and response information
func requestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if detailed logging is enabled
		if strings.ToUpper(os.Getenv("ENABLE_REQUEST_LOGGING")) != "ON" {
			next.ServeHTTP(w, r)
			return
		}

		// Log request details
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		log.Printf("Headers: %v", r.Header)

		// Log request body if present
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil {
				log.Printf("Request Body: %s", string(bodyBytes))
				// Restore the body for further processing
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// Create a response wrapper to capture the response
		rw := &responseWriter{
			ResponseWriter: w,
			body:          new(bytes.Buffer),
		}

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Log response details
		log.Printf("Response Status: %d", rw.status)
		log.Printf("Response Headers: %v", w.Header())
		log.Printf("Response Body: %s", rw.body.String())
	})
}

// responseWriter is a custom response writer that captures the response
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

func main() {
	// Load environment variables
	log.Printf("Loading env variables") 

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize router
	r := mux.NewRouter()
	log.Printf("Router initialized")

	// Apply middlewares
		r.Use(requestLoggerMiddleware)


	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		AllowedMethods:   strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ","),
		AllowedHeaders:   strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ","),
		AllowCredentials: strings.ToLower(os.Getenv("CORS_ALLOW_CREDENTIALS")) == "true",
		Logger: log.New(os.Stdout, "CORS: ", log.LstdFlags),
	})
	handler := c.Handler(r)

	// Initialize database
	db := initDB()
	log.Printf("Database connected")

	// Initialize handlers
	h := NewHandler(db)

	// Routes
	r.HandleFunc("/api/products", h.GetProducts).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.GetProduct).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products", h.CreateProduct).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.UpdateProduct).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/cart", h.GetCart).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/cart/add", h.AddToCart).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/cart/remove/{id}", h.RemoveFromCart).Methods("DELETE", "OPTIONS")

	// Start server
	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(log.Writer(), handler)))
	log.Printf("Server started")
	defer db.Close()
} 
