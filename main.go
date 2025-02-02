package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	log.Printf("Loading env variables") 

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize router
	r := mux.NewRouter()
	log.Printf("Router initialized")

	// Initialize database
	db := initDB()
	log.Printf("Database connected")
	defer db.Close()

	// Initialize handlers
	h := NewHandler(db)

	// Routes
	r.HandleFunc("/api/products", h.GetProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/api/products", h.CreateProduct).Methods("POST")
	r.HandleFunc("/api/cart", h.GetCart).Methods("GET")
	r.HandleFunc("/api/cart/add", h.AddToCart).Methods("POST")
	r.HandleFunc("/api/cart/remove/{id}", h.RemoveFromCart).Methods("DELETE")

	// Start server
	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
} 