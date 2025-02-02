package main

import (
	"kids-shop/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Loading env variables") 

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	db := initDB()
	log.Printf("Database connected")

	h := NewHandler(db)
	
	r := setupRouter(h)

	// Apply CORS middleware
	handler := middleware.NewCORS()(r)

	log.Printf("Router initialized")

	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(log.Writer(), handler)))
	log.Printf("Server started")
	defer db.Close()
} 
