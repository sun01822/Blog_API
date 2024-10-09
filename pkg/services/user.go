package services

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/types"
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
	if err := svc.repo.LoginRepo(email, password); err != nil {
		return errors.New("Log in Failed")
	}
	return nil
}

// CreateUser implements domain.UserService.
func (svc *userService) CreateUser(reqUser *types.SignUpRequest) error {
	user := &models.User{
		Email:          reqUser.Email,
		Password:       reqUser.Password,
		Gender:         reqUser.Gender,
		DateOfBirth:    reqUser.DateOfBirth,
		Job:            reqUser.Job,
		City:           reqUser.City,
		ZipCode:        reqUser.ZipCode,
		ProfilePicture: reqUser.ProfilePicture,
		FirstName:      reqUser.FirstName,
		LastName:       reqUser.LastName,
		Phone:          reqUser.Phone,
		Street:         reqUser.Street,
		State:          reqUser.State,
		Country:        reqUser.Country,
		Latitude:       reqUser.Latitude,
		Longitude:      reqUser.Longitude,
	}

	if err := svc.repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

// DeleteUser implements domain.UserService.
func (svc *userService) DeleteUser(id uint) error {
	if err := svc.repo.DeleteUserRepo(id); err != nil {
		return err
	}
	return nil
}

// GetUser implements domain.UserService.
func (svc *userService) GetUser(id uint) (models.User, error) {
	user, err := svc.repo.GetUserRepo(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetUsers implements domain.UserService.
func (svc *userService) GetUsers(pagination *utils.Page) ([]models.User, error) {
	users, err := svc.repo.GetUsersRepo(pagination)
	if err != nil {
		return users, err
	}
	return users, nil
}

// UpdateUser implements domain.UserService.
func (svc *userService) UpdateUser(user *models.User) error {
	if err := svc.repo.UpdateUserRepo(user); err != nil {
		return err
	}
	return nil
}
