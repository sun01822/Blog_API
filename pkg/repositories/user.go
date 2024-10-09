package repositories

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

// Parent struct to implement interface binding
type userRepo struct {
	d *gorm.DB
}

// Interface binding
func NewUserRepo(db *gorm.DB) domain.UserRepository {
	return &userRepo{
		d: db,
	}
}

// Login implements domain.UserRepository.
func (repo *userRepo) LoginRepo(email string, password string) error {
	// Find the user by user_name
	var existingUser models.User
	if err := repo.d.Where("email = ?", email).First(&existingUser).Error; err != nil {
		return err
	}
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err := utils.ComparePassword(existingUser.Password, password); err != nil {
		return err
	}
	// Otherwise, we are good to go, so return a nil error.
	return nil
}

// CreateUser implements domain.UserRepository.
func (repo *userRepo) CreateUser(user *models.User) error {
	userEmail := user.Email
	var existingUser models.User

	existingErr := repo.d.Where("email = ?", userEmail).First(&existingUser).Error
	if existingErr == nil {
		return errors.New("user already exists with same email")
	}

	// Hash the user password
	user.Password = utils.HashPassword(user.Password)

	createErr := repo.d.Create(user).Error
	if createErr != nil {
		return createErr
	}
	return nil
}

// DeleteUser implements domain.UserRepository.
func (repo *userRepo) DeleteUserRepo(id uint) error {
	var user models.User
	err := repo.d.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUser implements domain.UserRepository.
func (repo *userRepo) GetUser(id uint) (models.User, error) {
	var user models.User
	respErr := repo.d.Where("id = ?", id).First(&user).Error
	if respErr != nil {
		return user, respErr
	}
	return user, nil
}

// GetUsers implements domain.UserRepository.
func (repo *userRepo) GetUsersRepo(pagination *utils.Page) ([]models.User, error) {
	var users []models.User
	err := repo.d.Offset(*pagination.Offset).Limit(*pagination.Limit).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

// UpdateUser implements domain.UserRepository.
func (repo *userRepo) UpdateUserRepo(user *models.User) error {
	err := repo.d.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
