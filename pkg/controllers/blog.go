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
	tempID := c.Param("userID")
	userID, err := strconv.ParseUint(tempID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	reqBlogPost := &types.BlogPostRequest{}
	if err := c.Bind(reqBlogPost); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := reqBlogPost.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	currentTime := time.Now()
	blog := &models.BlogPost{
		UserID:       uint(userID),
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
	return c.JSON(http.StatusCreated, "Blog post created successfully")
}

// GetBlogPost implements domain.BlogController.
func (ctr *blogController) GetBlogPost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("postID"), 10, 64)
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
	blogPosts, err := ctr.svc.GetBlogPosts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, blogPosts)
}


// GetBlogPosts implements domain.BlogController.
func (ctr *blogController) GetBlogPostsOfUser(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	blogPosts, err := ctr.svc.GetBlogPostsOfUser(uint(userID))
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
	tempUserID := c.Param("userID")
	userID, err := strconv.ParseUint(tempUserID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	tempID := c.Param("postID")
	id, err := strconv.ParseUint(tempID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Blog post not found")
	}
	if existingBlogPost.UserID != uint(userID) {
		return c.JSON(http.StatusUnauthorized, "You are not authorized to update this blog post")
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
		UserID: uint(userID),
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
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	id, err := strconv.ParseUint(c.Param("postID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Blog post not found")
	}
	if existingBlogPost.UserID != uint(userID) {
		return c.JSON(http.StatusUnauthorized, "You are not authorized to delete this blog post")
	}
	if err := ctr.svc.DeleteBlogPost(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Blog post deleted successfully")
}

// AddAndRemoveLike implements domain.BlogController.
func (ctr *blogController) AddAndRemoveLike(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	tempID := c.Param("postID")
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
	if err := ctr.svc.AddAndRemoveLike(&existingBlogPost, uint(userID)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Like added/removed successfully")
}

// AddComment implements domain.BlogController.
func (ctr *blogController) AddComment(c echo.Context) error {
	reqComment := &types.Comment{}
	if err := c.Bind(reqComment); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := reqComment.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	tempID := c.Param("postID")
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
	comment := &models.Comment{
		Content:  reqComment.Content,
		UserID:   uint(userID),
		BlogPostID: existingBlogPost.ID,
		BlogPost: existingBlogPost,
	}
	if err := ctr.svc.AddComment(&existingBlogPost, comment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, comment)
}

// GetCommentByUserID implements domain.BlogController.
func (ctr *blogController) GetCommentByUserID(c echo.Context) error {
	tempID := c.Param("postID")
	id, err := strconv.ParseUint(tempID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	commentID, err := strconv.ParseUint(c.Param("commentID"), 10, 64)
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
	comment, err := ctr.svc.GetCommentByUserID(&existingBlogPost, uint(commentID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, comment)
}


// GetComments implements domain.BlogController.
func (ctr *blogController) GetComments(c echo.Context) error {
	tempID := c.Param("postID")
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
	comments, err := ctr.svc.GetComments(&existingBlogPost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, comments)
}

// DeleteComment implements domain.BlogController.
func (ctr *blogController) DeleteComment(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	tempID := c.Param("postID")
	id, err := strconv.ParseUint(tempID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	commentID, err := strconv.ParseUint(c.Param("commentID"), 10, 64)
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
	if _, err := ctr.svc.GetCommentByUserID(&existingBlogPost, uint(userID)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ctr.svc.DeleteComment(&existingBlogPost, uint(commentID)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Comment deleted successfully")
}

// UpdateComment implements domain.BlogController.
func (ctr *blogController) UpdateComment(c echo.Context) error {
	reqComment := &types.Comment{}
	if err := c.Bind(reqComment); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := reqComment.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tempID := c.Param("postID")
	id, err := strconv.ParseUint(tempID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	commentID, err := strconv.ParseUint(c.Param("commentID"), 10, 64)
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
	if _, err := ctr.svc.GetCommentByUserID(&existingBlogPost, uint(userID)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	comment := &models.Comment{
		Model: gorm.Model{ID: uint(commentID)},
		Content:  reqComment.Content,
	}
	if err := ctr.svc.UpdateComment(&existingBlogPost, comment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Comment updated successfully")
}
