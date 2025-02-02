package main

import (
	"kids-shop/middleware"

	"github.com/gorilla/mux"
)

func setupRouter(h *Handler) (*mux.Router) {
	r := mux.NewRouter()

	// Apply middlewares
	r.Use(middleware.RequestLogger)
	
	// Routes
	r.HandleFunc("/api/products", h.GetProducts).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.GetProduct).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products", h.CreateProduct).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.UpdateProduct).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.DeleteProduct).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/api/cart", h.GetCart).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/cart/add", h.AddToCart).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/cart/remove/{id}", h.RemoveFromCart).Methods("DELETE", "OPTIONS")

	return r
} 