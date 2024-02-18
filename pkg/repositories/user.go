package repositories

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
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


// CreateUser implements domain.UserRepository.
func (repo *userRepo) CreateUser(user *models.User) error {
	err := repo.d.Create(user).Error
	if err != nil {
		return err
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
