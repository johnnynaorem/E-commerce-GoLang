package services

import (
	"e-commerce/internal/domain/models"
	"e-commerce/internal/domain/repositories"
)

type ProductService struct {
	ProductRepo repositories.ProductRepository
}

func NewProductService(prodRepo repositories.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: prodRepo}
}

func (ps *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return ps.ProductRepo.CreateProduct(product)
}

func (ps *ProductService) UpdateProduct(product *models.Product) (*models.Product, error) {
	return ps.ProductRepo.Update(product)
}

func (ps *ProductService) DeleteProduct(productId uint) error {
	return ps.ProductRepo.DeleteProduct(productId)
}

func (ps *ProductService) FindProductById(productId string) (*models.Product, error) {
	return ps.ProductRepo.GetProductById(productId)
}

func (ps *ProductService) FindAllProduct() ([]models.Product, error) {
	return ps.ProductRepo.GetAllProduct()
}
