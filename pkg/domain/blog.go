package domain

import (
	"Blog_API/pkg/models"
	"Blog_API/pkg/types"
	"github.com/labstack/echo/v4"
)

// For database UserRepository opearation (call from service)
type BlogRepository interface {
	CreateBlogPost(blogPost models.BlogPost) error
	GetBlogPost(blogID string) (models.BlogPost, error)
	GetBlogPosts() ([]models.BlogPost, error)
	GetBlogPostsOfUser(userID string) ([]models.BlogPost, error)
	//UpdateBlogPostRepo(blogPost *models.BlogPost) error
	//DeleteBlogPostRepo(id uint) error
	//AddAndRemoveLikeRepo(blogPost *models.BlogPost, userID uint) (string, error)
	//AddCommentRepo(blogPost *models.BlogPost, comment *models.Comment) error
	//GetCommentByUserIDRepo(blogPost *models.BlogPost, commentID uint) (models.Comment, error)
	//GetCommentsRepo(blogPost *models.BlogPost) ([]models.Comment, error)
	//DeleteCommentRepo(blogPost *models.BlogPost, commentID uint) error
	//UpdateCommentRepo(blogPost *models.BlogPost, comment *models.Comment) error
}

// For service operation (call from controller)
type BlogService interface {
	CreateBlogPost(reqBlogPost types.BlogPostRequest, userID string) (types.BlogResp, error)
	GetBlogPost(blogID string) (types.BlogResp, error)
	GetBlogPosts() ([]types.BlogResp, error)
	GetBlogPostsOfUser(userID string) ([]types.BlogResp, error)
	//UpdateBlogPost(blogPost *models.BlogPost) error
	//DeleteBlogPost(id uint) error
	//AddAndRemoveLike(blogPost *models.BlogPost, userID uint) (string, error)
	//AddComment(blogPost *models.BlogPost, comment *models.Comment) error
	//GetCommentByUserID(blogPost *models.BlogPost, commentID uint) (models.Comment, error)
	//GetComments(blogPost *models.BlogPost) ([]models.Comment, error)
	//DeleteComment(blogPost *models.BlogPost, commentID uint) error
	//UpdateComment(blogPost *models.BlogPost, comment *models.Comment) error
}

// For controller operation (call from main)
type BlogController interface {
	CreateBlogPost(c echo.Context) error
	GetBlogPost(c echo.Context) error
	GetBlogPosts(c echo.Context) error
	GetBlogPostsOfUser(c echo.Context) error
	//UpdateBlogPost(c echo.Context) error
	//DeleteBlogPost(c echo.Context) error
	//AddAndRemoveLike(c echo.Context) error
	//AddComment(c echo.Context) error
	//GetCommentByUserID(c echo.Context) error
	//GetComments(c echo.Context) error
	//DeleteComment(c echo.Context) error
	//UpdateComment(c echo.Context) error
}
