package service

import "kids-shop/internal/repository/postgres"

type Services struct {
	Product ProductService
	// Add other services as needed
}

func NewServices(repos *postgres.Repositories) *Services {
	return &Services{
		Product: NewProductService(repos.Product),
		// Initialize other services
	}
} 