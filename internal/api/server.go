package api

import (
	"database/sql"
	"kids-shop/config"
	"kids-shop/middleware"
	"log"
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
	// Setup router
	router := setupRouter( s.db)

	// Apply CORS middleware
	s.handler = middleware.NewCORS()(router)
}

func (s *Server) Start() error {
	log.Printf("Starting server on port %s", s.config.Server.Port)
	return http.ListenAndServe(":"+s.config.Server.Port, s.handler)
} 