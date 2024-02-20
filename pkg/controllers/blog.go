package controllers

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	blogPost := &models.BlogPost{}
	if err := c.Bind(blogPost); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := ctr.svc.CreateBlogPost(blogPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, blogPost)
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
	blogPost := &models.BlogPost{}
	if err := c.Bind(blogPost); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := ctr.svc.UpdateBlogPost(blogPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, blogPost)
}

// DeleteBlogPost implements domain.BlogController.
func (ctr *blogController) DeleteBlogPost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}
	if err := ctr.svc.DeleteBlogPost(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Blog post deleted successfully")
}

