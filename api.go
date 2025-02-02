package main

import (
	"kids-shop/middleware"

	"github.com/gorilla/mux"
)

func setupRouter(h *Handler) (*mux.Router) {
	r := mux.NewRouter()

	// Apply middlewares
	r.Use(middleware.RequestLogger)
	
	// Auth routes
	r.HandleFunc("/api/auth/login", h.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/auth/register", h.Register).Methods("POST", "OPTIONS")

	// User profile routes
	r.HandleFunc("/api/profile", h.GetUserProfile).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/profile", h.UpdateUserProfile).Methods("PUT", "OPTIONS")

	// Cart routes
	r.HandleFunc("/api/cart", h.GetCart).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/cart", h.AddToCart).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/cart/{id}", h.UpdateCartItem).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/cart/{id}", h.RemoveFromCart).Methods("DELETE", "OPTIONS")

	// Order routes
	r.HandleFunc("/api/orders", h.GetOrders).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/orders", h.CreateOrder).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/orders/{id}", h.GetOrder).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/orders/{id}", h.UpdateOrder).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/orders/{id}", h.DeleteOrder).Methods("DELETE", "OPTIONS")

	// Product routes (existing)
	r.HandleFunc("/api/products", h.GetProducts).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.GetProduct).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products", h.CreateProduct).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.UpdateProduct).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/products/{id}", h.DeleteProduct).Methods("DELETE", "OPTIONS")

	return r
} 