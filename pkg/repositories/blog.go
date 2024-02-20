package repositories

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"gorm.io/gorm"
)

// Parent struct to implement interface binding
type blogRepo struct {
	d *gorm.DB
}

// Interface binding
func NewBlogRepo(db *gorm.DB) domain.BlogRepository {
	return &blogRepo{
		d: db,
	}
}

// CreateBlogPost implements domain.BlogRepository.
func (repo *blogRepo) CreateBlogPost(blogPost *models.BlogPost) error {
	err := repo.d.Create(blogPost).Error
	if err != nil {
		return err
	}
	return nil
}

// GetBlogPost implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPost(id uint) (models.BlogPost, error) {
	var blogPost models.BlogPost
	err := repo.d.Where("id = ?", id).First(&blogPost).Error
	if err != nil {
		return blogPost, err
	}
	return blogPost, nil
}

// GetBlogPosts implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPosts(userID uint) ([]models.BlogPost, error) {
	var blogPosts []models.BlogPost
	err := repo.d.Where("user_id = ?", userID).Find(&blogPosts).Error
	if err != nil {
		return blogPosts, err
	}
	return blogPosts, nil
}

// UpdateBlogPost implements domain.BlogRepository.
func (repo *blogRepo) UpdateBlogPost(blogPost *models.BlogPost) error {
	err := repo.d.Save(blogPost).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteBlogPost implements domain.BlogRepository.
func (repo *blogRepo) DeleteBlogPost(id uint) error {
	err := repo.d.Where("id = ?", id).Delete(&models.BlogPost{}).Error
	if err != nil {
		return err
	}
	return nil
}
