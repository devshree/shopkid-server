package service

import (
	"kids-shop/internal/domain/models"
	"kids-shop/internal/domain/services"
)

type ProductService interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id int) error
}

type productService struct {
	repo services.ProductRepository
}

func NewProductService(repo services.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *productService) Delete(id int) error {
	return s.repo.Delete(id)
}

// ... implement other methods 