package domain

import (
	"Blog_API/pkg/models"
	"github.com/labstack/echo/v4"
)

// For database Repository opearation (call from service)
type BlogRepository interface {
	CreateBlogPost(blogPost *models.BlogPost) error
	GetBlogPost(id uint) (models.BlogPost, error)
	GetBlogPosts(userID uint) ([]models.BlogPost, error)
	UpdateBlogPost(blogPost *models.BlogPost) error
	DeleteBlogPost(id uint) error
}

// For service operation (call from controller)
type BlogService interface {
	CreateBlogPost(blogPost *models.BlogPost) error
	GetBlogPost(id uint) (models.BlogPost, error)
	GetBlogPosts(userID uint) ([]models.BlogPost, error)
	UpdateBlogPost(blogPost *models.BlogPost) error
	DeleteBlogPost(id uint) error
}

// For controller operation (call from main)
type BlogController interface {
	CreateBlogPost(c echo.Context) error
	GetBlogPost(c echo.Context) error
	GetBlogPosts(c echo.Context) error
	UpdateBlogPost(c echo.Context) error
	DeleteBlogPost(c echo.Context) error
}


