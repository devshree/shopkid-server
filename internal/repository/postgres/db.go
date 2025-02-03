package postgres

import (
	"database/sql"
	"fmt"
	"kids-shop/config"
)

func NewDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

type Repositories struct {
	Product *ProductRepository
	// Add other repositories as needed
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Product: NewProductRepository(db),
		// Initialize other repositories
	}
} 