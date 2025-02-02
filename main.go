package main

import (
	"log"
	"net/http"

	"kids-shop/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Loading env variables") 

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	r := mux.NewRouter()
	log.Printf("Router initialized")

	// Apply middlewares
	r.Use(middleware.RequestLogger)
	
	// Apply CORS middleware
	handler := middleware.NewCORS()(r)

	db := initDB()
	log.Printf("Database connected")

	h := NewHandler(db)

	// Routes
	r.HandleFunc("/api/products", h.GetProducts).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.GetProduct).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products", h.CreateProduct).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.UpdateProduct).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.DeleteProduct).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/cart", h.GetCart).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/cart/add", h.AddToCart).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/cart/remove/{id}", h.RemoveFromCart).Methods("DELETE", "OPTIONS")

	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(log.Writer(), handler)))
	log.Printf("Server started")
	defer db.Close()
} 
