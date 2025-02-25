package persistence

import (
	"e-commerce/internal/domain/models"
	"e-commerce/internal/domain/repositories"
	"fmt"

	"gorm.io/gorm"
)

type UserRepositoryImp struct {
	DB *gorm.DB
}

func NewUserRepositoryImp(db *gorm.DB) repositories.UserRepository {
	return &UserRepositoryImp{DB: db}
}

func (ur *UserRepositoryImp) Create(user *models.User) (*models.User, error) {
	if err := ur.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepositoryImp) FindAll() ([]models.User, error) {
	var users []models.User
	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (ur *UserRepositoryImp) Update(user *models.User) (*models.User, error) {
	var existUser models.User
	if err := ur.DB.Where("email= ?", user.Email).First(&existUser).Error; err != nil {
		fmt.Println(user)
		return nil, err
	}
	var updateUser = user
	if err := ur.DB.Model(&models.User{}).Where("email= ?", user.Email).Updates(&updateUser).Error; err != nil {
		return nil, err
	}
	return updateUser, nil
}
func (ur *UserRepositoryImp) Delete(user *models.User) error {
	var email = user.Email
	if err := ur.DB.Where("email = ?", email).Error; err != nil {
		return fmt.Errorf("user not found")
	}
	if err := ur.DB.Where("email = ?", email).Delete(user).Error; err != nil {
		return fmt.Errorf("failed to delete user")
	}
	return nil
}

func (ur *UserRepositoryImp) FindByID(email string) (*models.User, error) {
	var user models.User
	if err := ur.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
