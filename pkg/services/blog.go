package services

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
)

// Parent struct to implement interface binding
type blogService struct {
	repo domain.BlogRepository
}

// Interface binding
func NewBlogService(repo domain.BlogRepository) domain.BlogService {
	return &blogService{
		repo: repo,
	}
}

// CreateBlogPost implements domain.BlogService.
func (svc *blogService) CreateBlogPost(blogPost *models.BlogPost) error {
	if err := svc.repo.CreateBlogPost(blogPost); err != nil {
		return err
	}
	return nil
}

// GetBlogPost implements domain.BlogService.
func (svc *blogService) GetBlogPost(id uint) (models.BlogPost, error) {
	blogPost, err := svc.repo.GetBlogPost(id)
	if err != nil {
		return blogPost, err
	}
	return blogPost, nil
}

// GetBlogPosts implements domain.BlogService.
func (svc *blogService) GetBlogPosts(userID uint) ([]models.BlogPost, error) {
	blogPosts, err := svc.repo.GetBlogPosts(userID)
	if err != nil {
		return blogPosts, err
	}
	return blogPosts, nil
}

// UpdateBlogPost implements domain.BlogService.
func (svc *blogService) UpdateBlogPost(blogPost *models.BlogPost) error {
	if err := svc.repo.UpdateBlogPost(blogPost); err != nil {
		return err
	}
	return nil
}

// DeleteBlogPost implements domain.BlogService.
func (svc *blogService) DeleteBlogPost(id uint) error {
	if err := svc.repo.DeleteBlogPost(id); err != nil {
		return err
	}
	return nil
}
