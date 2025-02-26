package persistence

import (
	"e-commerce/internal/domain/models"
	"e-commerce/internal/domain/repositories"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) repositories.ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}

func (pr *ProductRepositoryImpl) CreateProduct(product *models.Product) (*models.Product, error) {
	if error := pr.DB.Create(product).Error; error != nil {
		return nil, error
	}
	return product, nil
}

func (pr *ProductRepositoryImpl) GetAllProduct() ([]models.Product, error) {
	var products []models.Product
	if error := pr.DB.Find(&products).Error; error != nil {
		return nil, error
	}
	return products, nil
}

func (pr *ProductRepositoryImpl) DeleteProduct(productId uint) error {
	var product models.Product
	if err := pr.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		return fmt.Errorf("product not found")
	}
	if err := pr.DB.Where("id = ?", productId).Delete(product).Error; err != nil {
		return fmt.Errorf("failed to delete product")
	}
	return nil
}

func (ur *ProductRepositoryImpl) GetProductById(id string) (*models.Product, error) {
	var product models.Product
	if err := ur.DB.Where("ID = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (ur *ProductRepositoryImpl) Update(product *models.Product) (*models.Product, error) {
	var existProduct models.Product
	if err := ur.DB.Where("id= ?", product.ID).First(&existProduct).Error; err != nil {
		return nil, err
	}
	var updateUser = product
	if err := ur.DB.Model(&models.User{}).Where("ID= ?", product.ID).Updates(&updateUser).Error; err != nil {
		return nil, err
	}
	return updateUser, nil
}
