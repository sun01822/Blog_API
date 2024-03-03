package services

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/utils"
	"errors"
)

// Parent struct to implement interface binding
type userService struct {
	repo domain.UserRepository
}

// Interface binding
func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{
		repo: repo,
	}
}

// Login implements domain.UserService.
func (svc *userService) Login(email string, password string) error {
	if err := svc.repo.Login(email, password); err != nil {
		return errors.New("Log in Failed")
	}
	return nil
}

// CreateUser implements domain.UserService.
func (svc *userService) CreateUser(user *models.User) error {
	if err := svc.repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

// DeleteUser implements domain.UserService.
func (svc *userService) DeleteUser(id uint) error {
	if err := svc.repo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}

// GetUser implements domain.UserService.
func (svc *userService) GetUser(id uint) (models.User, error) {
	user, err := svc.repo.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetUsers implements domain.UserService.
func (svc *userService) GetUsers(pagination *utils.Page) ([]models.User, error) {
	users, err := svc.repo.GetUsers(pagination)
	if err != nil {
		return users, err
	}
	return users, nil
}

// UpdateUser implements domain.UserService.
func (svc *userService) UpdateUser(user *models.User) error {
	if err := svc.repo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}
