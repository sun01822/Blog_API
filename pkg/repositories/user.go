package repositories

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"gorm.io/gorm"
	"Blog_API/pkg/utils"
	"errors"
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
func (repo *userRepo) Login(email string, password string) error {
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
	err := repo.d.Where("email = ?", userEmail).First(&existingUser).Error
	if err == nil {
		return errors.New("User already exists with same email")
	}

	// Hash the user password
	user.Password = utils.HashPassword(user.Password)

	err2 := repo.d.Create(user).Error
	if err2 != nil {
		return err2
	}
	return nil
}

// DeleteUser implements domain.UserRepository.
func (repo *userRepo) DeleteUser(id int) error {
	var user models.User
	err := repo.d.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUser implements domain.UserRepository.
func (repo *userRepo) GetUser(id int) (models.User, error) {
	var user models.User
	err := repo.d.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetUsers implements domain.UserRepository.
func (repo *userRepo) GetUsers() ([]models.User, error) {
	var users []models.User
	err := repo.d.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

// UpdateUser implements domain.UserRepository.
func (repo *userRepo) UpdateUser(user *models.User) error {
	err := repo.d.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
