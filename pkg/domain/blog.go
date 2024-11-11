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
	GetBlogPostsOfUser(userID string, blogIDs []string) ([]models.BlogPost, error)
	UpdateBlogPost(blogPost models.BlogPost) error
	DeleteBlogPost(blogID string) error
	AddAndRemoveReaction(userID string, reactionID uint64, blogPost models.BlogPost) (models.BlogPost, error)
	AddComment(blogPost models.BlogPost, comment models.Comment) (models.BlogPost, error)
	GetComments(blogID string, commentIDs []string) ([]models.Comment, error)
	//DeleteCommentRepo(blogPost *models.BlogPost, commentID uint) error
	//UpdateCommentRepo(blogPost *models.BlogPost, comment *models.Comment) error
}

// For service operation (call from controller)
type BlogService interface {
	CreateBlogPost(reqBlogPost types.BlogPostRequest, userID string) (types.BlogResp, error)
	GetBlogPost(blogID string) (types.BlogResp, error)
	GetBlogPosts() ([]types.BlogResp, error)
	GetBlogPostsOfUser(userID string, blogIDs []string) ([]types.BlogResp, error)
	UpdateBlogPost(userID string, blogID string, blogPost types.UpdateBlogPostRequest) (types.BlogResp, error)
	DeleteBlogPost(userID string, blogID string) error
	AddAndRemoveReaction(userID string, blogID string, reactionID uint64) (types.BlogResp, error)
	AddComment(userID string, blogID string, comment types.Comment) (types.BlogResp, error)
	GetComments(userID string, blogID string, commentIDs []string) ([]types.CommentResp, error)
	//DeleteComment(blogPost *models.BlogPost, commentID uint) error
	//UpdateComment(blogPost *models.BlogPost, comment *models.Comment) error
}

// For controller operation (call from main)
type BlogController interface {
	CreateBlogPost(c echo.Context) error
	GetBlogPost(c echo.Context) error
	GetBlogPosts(c echo.Context) error
	GetBlogPostsOfUser(c echo.Context) error
	UpdateBlogPost(c echo.Context) error
	DeleteBlogPost(c echo.Context) error
	AddAndRemoveReaction(c echo.Context) error
	AddComment(c echo.Context) error
	GetComments(c echo.Context) error
	//DeleteComment(c echo.Context) error
	//UpdateComment(c echo.Context) error
}
