package services

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
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


// CreateUser implements domain.UserService.
func (svc *userService) CreateUser(user *models.User) error {
	if err := svc.repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

// DeleteUser implements domain.UserService.
func (svc *userService) DeleteUser(id int) error {
	if err := svc.repo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}

// GetUser implements domain.UserService.
func (svc *userService) GetUser(id int) (models.User, error) {
	user, err := svc.repo.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetUsers implements domain.UserService.
func (svc *userService) GetUsers() ([]models.User, error) {
	users, err := svc.repo.GetUsers()
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
