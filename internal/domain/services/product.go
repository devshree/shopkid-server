package services

import "kids-shop/internal/domain/models"

type ProductRepository interface {
    GetAll() ([]models.Product, error)
    GetByID(id int) (*models.Product, error)
    Create(product *models.Product) error
    Update(product *models.Product) error
    Delete(id int) error
} 