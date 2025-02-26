package repositories

import "e-commerce/internal/domain/models"

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetAllProduct() ([]models.Product, error)
	GetProductById(id string) (*models.Product, error)
	DeleteProduct(id uint) error
	Update(prouct *models.Product) (*models.Product, error)
}
