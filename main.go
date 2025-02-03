package main

import (
	"kids-shop/middleware"
	"log"
	"net/http"

	"kids-shop/postgres"
	"kids-shop/service"

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

	// Initialize repositories
	productRepo := postgres.NewProductRepository(db)
	
	// Initialize services
	productService := service.NewProductService(productRepo)
	
	// Initialize handler with services
	handler := NewHandler(db, productService)
	
	// Setup router
	router := setupRouter(handler)

	// Apply CORS middleware
	handler = middleware.NewCORS()(router)

	log.Printf("Router initialized")

	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(log.Writer(), handler)))
	log.Printf("Server started")
	defer db.Close()
} 
