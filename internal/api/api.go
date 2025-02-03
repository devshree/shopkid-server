package api

import (
	"database/sql"
	"kids-shop/internal/api/handlers"
	"kids-shop/middleware"

	"github.com/gorilla/mux"
)

func setupRouter( db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Apply middlewares
	r.Use(middleware.RequestLogger)
	r.Use(middleware.AuthMiddleware)
	
	// Auth routes
	ah := handlers.NewAuthHandler(db)
	r.HandleFunc("/api/auth/login", ah.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/auth/register", ah.Register).Methods("POST", "OPTIONS")

	// User profile routes
	uh := handlers.NewUserHandler(db)
	r.HandleFunc("/api/profile", uh.GetUserProfile).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/profile", uh.UpdateUserProfile).Methods("PUT", "OPTIONS")

	// Cart routes
	ch := handlers.NewCartHandler(db)
	r.HandleFunc("/api/cart", ch.GetCart).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/cart", ch.AddToCart).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/cart/{id}", ch.UpdateCartItem).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/cart/{id}", ch.RemoveFromCart).Methods("DELETE", "OPTIONS")

	// Order routes
	oh := handlers.NewOrderHandler(db)
	r.HandleFunc("/api/orders", oh.GetOrders).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/orders", oh.CreateOrder).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/orders/{id}", oh.GetOrder).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/orders/{id}", oh.UpdateOrder).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/orders/{id}", oh.DeleteOrder).Methods("DELETE", "OPTIONS")

	// Product routes
	productHandler := handlers.NewProductHandler(db)
	r.HandleFunc("/api/products", productHandler.GetProducts).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products/{id}", productHandler.GetProduct).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/products", productHandler.CreateProduct).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/products/{id}", productHandler.UpdateProduct).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/products/{id}", productHandler.DeleteProduct).Methods("DELETE", "OPTIONS")

	return r
} 