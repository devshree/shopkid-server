package main

import (
	"kids-shop/config"
	"kids-shop/internal/api"
	"kids-shop/internal/repository/postgres"
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

	// Initialize and start HTTP server
	server := api.NewServer(cfg, db)
	log.Fatal(server.Start())
} 