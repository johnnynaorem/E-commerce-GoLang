package services

import (
	"e-commerce/internal/domain/models"
	"e-commerce/internal/domain/repositories"
)

type UserService struct {
	UserRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	return us.UserRepo.Create(user)
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	return us.UserRepo.FindAll()
}
func (us *UserService) Update(user *models.User) (*models.User, error) {
	return us.UserRepo.Update(user)
}
func (us *UserService) Delete(user *models.User) error {
	return us.UserRepo.Delete(user)
}
func (us *UserService) FindByID(email string) (*models.User, error) {
	return us.UserRepo.FindByID(email)
}
