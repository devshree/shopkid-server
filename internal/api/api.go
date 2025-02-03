package api

import (
	"database/sql"
	"kids-shop/internal/api/handlers"
	"kids-shop/middleware"

	"github.com/gorilla/mux"
)

func setupRouter(h *Handler, db *sql.DB) *mux.Router {
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

	// Product routes
	productHandler := handlers.NewProductHandler(db)
	r.HandleFunc("/api/products", productHandler.GetProducts).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products/{id}", productHandler.GetProduct).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products", productHandler.CreateProduct).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/products/{id}", productHandler.UpdateProduct).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/products/{id}", productHandler.DeleteProduct).Methods("DELETE", "OPTIONS")

	return r
} 