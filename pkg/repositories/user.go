package repositories

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/utils"
	userconsts "Blog_API/pkg/utils/consts/user"
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
func (repo *userRepo) Login(email string, password string) (string, error) {

	var existingUser models.User

	if err := repo.d.Where("email = ?", email).First(&existingUser).Error; err != nil {
		return "", err
	}

	if err := utils.ComparePassword(existingUser.Password, password); err != nil {
		return "", err
	}

	return existingUser.ID, nil
}

// CreateUser implements domain.UserRepository.
func (repo *userRepo) CreateUser(user models.User) error {

	var existingUser models.User
	userEmail := user.Email

	err := repo.d.Where("email = ?", userEmail).First(&existingUser).Error
	if err == nil {
		return errors.New(userconsts.UserEmailAlreadyExists)
	}

	// Hash the user password
	user.Password = utils.HashPassword(user.Password)

	createErr := repo.d.Create(&user).Error
	if createErr != nil {
		return createErr
	}

	return nil
}

//// DeleteUser implements domain.UserRepository.
//func (repo *userRepo) DeleteUserRepo(id uint) error {
//	var user models.User
//	err := repo.d.Where("id = ?", id).Delete(&user).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
//

// GetUser implements domain.UserRepository.
func (repo *userRepo) GetUser(userID string) (models.User, error) {

	var user models.User

	err := repo.d.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUsers implements domain.UserRepository.
func (repo *userRepo) GetUsers(pagination utils.Page) ([]models.User, error) {

	var users []models.User

	query := repo.d.Model(&models.User{})

	if pagination.Offset > 0 {
		query = query.Offset(pagination.Offset)
	}

	if pagination.Limit > 0 {
		query = query.Limit(pagination.Limit)
	}

	err := query.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

// UpdateUser implements domain.UserRepository.
func (repo *userRepo) UpdateUser(user models.User) error {

	err := repo.d.Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}
