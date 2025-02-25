package repositories

import "e-commerce/internal/domain/models"

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(user *models.User) error
	FindByID(email string) (*models.User, error)
}
