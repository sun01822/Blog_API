package controllers

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/types"
	"net/http"
	"strconv"
	"time"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Parent struct to implement interface binding
type blogController struct {
	svc domain.BlogService
}

// Interface binding
func NewBlogController(svc domain.BlogService) domain.BlogController {
	return &blogController{
		svc: svc,
	}
}

// CreateBlogPost implements domain.BlogController.
func (ctr *blogController) CreateBlogPost(c echo.Context) error {
	reqBlogPost := &types.BlogPostRequest{}
	if err := c.Bind(reqBlogPost); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := reqBlogPost.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	currentTime := time.Now()
	blog := &models.BlogPost{
		UserID:       reqBlogPost.UserID,
		Title:        reqBlogPost.Title,
		ContentText:  reqBlogPost.ContentText,
		PhotoURL:     reqBlogPost.PhotoURL,
		Description:  reqBlogPost.Description,
		Category:     reqBlogPost.Category,
		IsPublished:  reqBlogPost.IsPublished,
		PublishedAt:  &currentTime,
	}
	if err := ctr.svc.CreateBlogPost(blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, blog)
}

// GetBlogPost implements domain.BlogController.
func (ctr *blogController) GetBlogPost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	blogPost, err := ctr.svc.GetBlogPost(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, blogPost)
}

// GetBlogPosts implements domain.BlogController.
func (ctr *blogController) GetBlogPosts(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	blogPosts, err := ctr.svc.GetBlogPosts(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, blogPosts)
}

// UpdateBlogPost implements domain.BlogController.
func (ctr *blogController) UpdateBlogPost(c echo.Context) error {
	blogPost := &types.UpdateBlogPostRequest{}
	if err := c.Bind(blogPost); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	tempID := c.Param("id")
	id, err := strconv.ParseUint(tempID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if existingBlogPost.ID == 0 {
		return c.JSON(http.StatusNotFound, "Blog post not found")
	}
	blog := &models.BlogPost{
		Model: gorm.Model{ID: uint(existingBlogPost.ID), CreatedAt: existingBlogPost.CreatedAt, UpdatedAt: time.Now(), DeletedAt: existingBlogPost.DeletedAt},
		Title:        blogPost.Title,
		ContentText:  blogPost.ContentText,
		PhotoURL:     blogPost.PhotoURL,
		Description:  blogPost.Description,
		Category:     blogPost.Category,
		IsPublished: existingBlogPost.IsPublished,
		PublishedAt: existingBlogPost.PublishedAt,
		Likes: existingBlogPost.Likes,
		UserID: existingBlogPost.UserID,
		LikesCount: existingBlogPost.LikesCount,
		Comments: existingBlogPost.Comments,
		CommentsCount: existingBlogPost.CommentsCount,
	}
	if blog.Title == "" {
		blog.Title = existingBlogPost.Title
	}
	if blog.ContentText == "" {
		blog.ContentText = existingBlogPost.ContentText
	}
	if blog.PhotoURL == "" {
		blog.PhotoURL = existingBlogPost.PhotoURL
	}
	if blog.Description == "" {
		blog.Description = existingBlogPost.Description
	}
	if blog.Category == "" {
		blog.Category = existingBlogPost.Category
	}
	if err := ctr.svc.UpdateBlogPost(blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Blog post updated successfully")
}

// DeleteBlogPost implements domain.BlogController.
func (ctr *blogController) DeleteBlogPost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if existingBlogPost.ID == 0 {
		return c.JSON(http.StatusNotFound, "Blog post not found")
	}
	if err := ctr.svc.DeleteBlogPost(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Blog post deleted successfully")
}

