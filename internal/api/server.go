package api

import (
	"database/sql"
	"kids-shop/config"
	"kids-shop/middleware"
	"net/http"
)

type Server struct {
	config   *config.Config
	handler   http.Handler 
	db       *sql.DB
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	server := &Server{
		config:   cfg,
		db:       db,
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	
	handler := NewHandler(s.db)
	
	// Setup router
	router := setupRouter(handler, s.db)

	// Apply CORS middleware
	s.handler = middleware.NewCORS()(router)
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.config.Server.Port, s.handler)
} 