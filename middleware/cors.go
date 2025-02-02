package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

// NewCORS creates a new CORS middleware handler
func NewCORS() func(http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		AllowedMethods:   strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ","),
		AllowedHeaders:   strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ","),
		AllowCredentials: strings.ToLower(os.Getenv("CORS_ALLOW_CREDENTIALS")) == "true",
	})
	return c.Handler
} 