package main

import (
	"kids-shop/config"
	"kids-shop/internal/api"
	"kids-shop/internal/repository/postgres"
	"kids-shop/internal/service"
	"log"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize DB
	db, err := postgres.NewDB(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repositories
	repos := postgres.NewRepositories(db)

	// Initialize services
	services := service.NewServices(repos)

	// Initialize and start HTTP server
	server := api.NewServer(cfg, services)
	log.Fatal(server.Start())
} 