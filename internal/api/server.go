package api

import (
	"kids-shop/config"
	"kids-shop/internal/api/handlers"
	"kids-shop/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	config   *config.Config
	router   *mux.Router
	services *service.Services
}

func NewServer(cfg *config.Config, services *service.Services) *Server {
	server := &Server{
		config:   cfg,
		router:   mux.NewRouter(),
		services: services,
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	productHandler := handlers.NewProductHandler(s.services.Product)

	api := s.router.PathPrefix("/api").Subrouter()
	
	// Product routes
	api.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	// Add other routes...
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.config.Server.Port, s.router)
} 