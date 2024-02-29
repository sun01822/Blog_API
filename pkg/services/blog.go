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
func (svc *blogService) GetBlogPosts() ([]models.BlogPost, error) {
	blogPosts, err := svc.repo.GetBlogPosts()
	if err != nil {
		return blogPosts, err
	}
	return blogPosts, nil
}

// GetBlogPosts implements domain.BlogService.
func (svc *blogService) GetBlogPostsOfUser(userID uint) ([]models.BlogPost, error) {
	blogPosts, err := svc.repo.GetBlogPostsOfUser(userID)
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

// AddAndRemoveLike implements domain.BlogService.
func (svc *blogService) AddAndRemoveLike(blogPost *models.BlogPost, userID uint) (string, error) {
	s, err := svc.repo.AddAndRemoveLike(blogPost, userID)
	if err != nil {
		return s, err
	}
	return s, nil
}

// AddComment implements domain.BlogService.
func (svc *blogService) AddComment(blogPost *models.BlogPost, comment *models.Comment) error {
	if err := svc.repo.AddComment(blogPost, comment); err != nil {
		return err
	}
	return nil
}

// GetCommentByUserID implements domain.BlogService.
func (svc *blogService) GetCommentByUserID(blogPost *models.BlogPost, commentID uint) (models.Comment, error) {
	comment, err := svc.repo.GetCommentByUserID(blogPost, commentID)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

// GetComments implements domain.BlogService.
func (svc *blogService) GetComments(blogPost *models.BlogPost) ([]models.Comment, error) {
	comments, err := svc.repo.GetComments(blogPost)
	if err != nil {
		return comments, err
	}
	return comments, nil
}

// DeleteComment implements domain.BlogService.
func (svc *blogService) DeleteComment(blogPost *models.BlogPost, commentID uint) error {
	if err := svc.repo.DeleteComment(blogPost, commentID); err != nil {
		return err
	}
	return nil
}

// UpdateComment implements domain.BlogService.
func (svc *blogService) UpdateComment(blogPost *models.BlogPost, comment *models.Comment) error {
	if err := svc.repo.UpdateComment(blogPost, comment); err != nil {
		return err
	}
	return nil
}
