package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get CORS settings from environment variables
		allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
		allowedMethods := os.Getenv("CORS_ALLOWED_METHODS")
		allowedHeaders := os.Getenv("CORS_ALLOWED_HEADERS")
		allowCredentials := os.Getenv("CORS_ALLOW_CREDENTIALS")
		maxAge := os.Getenv("CORS_MAX_AGE")

		// Handle multiple origins
		origin := r.Header.Get("Origin")
		if origin != "" {
			allowed := false
			for _, allowedOrigin := range allowedOrigins {
				if strings.TrimSpace(allowedOrigin) == origin {
					allowed = true
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
			// If no specific origin matches and "*" is in the allowed origins, allow all
			if !allowed {
				for _, allowedOrigin := range allowedOrigins {
					if strings.TrimSpace(allowedOrigin) == "*" {
						w.Header().Set("Access-Control-Allow-Origin", "*")
						break
					}
				}
			}
		}

		// Set other CORS headers
		if allowedMethods != "" {
			w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
		}
		if allowedHeaders != "" {
			w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		}
		if strings.ToLower(allowCredentials) == "true" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		if maxAge != "" {
			w.Header().Set("Access-Control-Max-Age", maxAge)
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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
	r.Use(corsMiddleware)
	r.Use(requestLoggerMiddleware)

	// Initialize database
	db := initDB()
	log.Printf("Database connected")
	defer db.Close()

	// Initialize handlers
	h := NewHandler(db)

	// Routes
	r.HandleFunc("/api/products", h.GetProducts).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.GetProduct).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products", h.CreateProduct).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/cart", h.GetCart).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/cart/add", h.AddToCart).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/cart/remove/{id}", h.RemoveFromCart).Methods("DELETE", "OPTIONS")

	// Start server
	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(log.Writer(), r)))
	log.Printf("Server started")
} 