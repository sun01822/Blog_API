package controllers

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/types"
	"Blog_API/pkg/utils/consts"
	blogconsts "Blog_API/pkg/utils/consts/blog"
	userconsts "Blog_API/pkg/utils/consts/user"
	"Blog_API/pkg/utils/response"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Parent struct to implement interface binding
type blogController struct {
	svc  domain.BlogService
	svc2 domain.Service
}

// Interface binding
func NewBlogController(svc domain.BlogService, svc2 domain.Service) domain.BlogController {
	return &blogController{
		svc:  svc,
		svc2: svc2,
	}
}

// CreateBlogPost implements domain.BlogController.
// @Summary Create a blog post
// @Description Create a blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param blogPost body types.BlogPostRequest true "Blog Post Request"
// @Success 200 {object} types.BlogResp "Blog Post Created Successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error creating blog"
// @Router /blog [post]
func (ctr *blogController) CreateBlogPost(c echo.Context) error {

	userID, parseErr := uuid.Parse(c.Get(userconsts.UserID).(string))
	if parseErr != nil {
		return response.ErrorResponse(c, parseErr, consts.InvalidDataRequest)
	}

	reqBlogPost := types.BlogPostRequest{}
	if bindErr := c.Bind(&reqBlogPost); bindErr != nil {
		return response.ErrorResponse(c, bindErr, consts.InvalidDataRequest)
	}

	if validationErr := reqBlogPost.Validate(); validationErr != nil {
		return response.ErrorResponse(c, validationErr, consts.ValidationError)
	}

	blog, err := ctr.svc.CreateBlogPost(reqBlogPost, userID.String())
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorCreatingBlog)
	}

	return response.SuccessResponse(c, blogconsts.BlogCreatedSuccessfully, blog)
}

// GetBlogPost implements domain.BlogController.
func (ctr *blogController) GetBlogPost(c echo.Context) error {

	reqBlogID := c.QueryParam(blogconsts.BlogID)

	if reqBlogID == "" {
		return response.ErrorResponse(c, errors.New(blogconsts.BlogIDRequired), consts.InvalidDataRequest)
	}

	blogPost, err := ctr.svc.GetBlogPost(reqBlogID)
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorGettingBlog)
	}

	return response.SuccessResponse(c, blogconsts.BlogFetchSuccessfully, blogPost)
}

// GetBlogPosts implements domain.BlogController.
// @Summary Get all blog posts
// @Description Get all blog posts
// @Tags Blog
// @Accept json
// @Produce json
// @Success 200 {array} types.BlogResp "Blogs Fetched Successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error getting blogs"
// @Router /blogs [get]
func (ctr *blogController) GetBlogPosts(c echo.Context) error {

	blogPosts, err := ctr.svc.GetBlogPosts()
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorGettingBlogs)
	}

	return response.SuccessResponse(c, blogconsts.BlogsFetchSuccessfully, blogPosts)
}

//// GetBlogPosts implements domain.BlogController.
//func (ctr *blogController) GetBlogPostsOfUser(c echo.Context) error {
//	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	blogPosts, err := ctr.svc.GetBlogPostsOfUser(uint(userID))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	return c.JSON(http.StatusOK, blogPosts)
//}
//
//// UpdateBlogPost implements domain.BlogController.
//func (ctr *blogController) UpdateBlogPost(c echo.Context) error {
//	blogPost := &types.UpdateBlogPostRequest{}
//	if err := c.Bind(blogPost); err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	tempUserID := c.Param("userID")
//	userID, err := strconv.ParseUint(tempUserID, 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	tempID := c.Param("postID")
//	id, err := strconv.ParseUint(tempID, 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Blog post not found")
//	}
//	if existingBlogPost.UserID != uint(userID) {
//		return c.JSON(http.StatusUnauthorized, "You are not authorized to update this blog post")
//	}
//	blog := &models.BlogPost{
//		Model:         gorm.Model{ID: uint(existingBlogPost.ID), CreatedAt: existingBlogPost.CreatedAt, UpdatedAt: time.Now(), DeletedAt: existingBlogPost.DeletedAt},
//		Title:         blogPost.Title,
//		ContentText:   blogPost.ContentText,
//		PhotoURL:      blogPost.PhotoURL,
//		Description:   blogPost.Description,
//		Category:      blogPost.Category,
//		IsPublished:   existingBlogPost.IsPublished,
//		PublishedAt:   existingBlogPost.PublishedAt,
//		Likes:         existingBlogPost.Likes,
//		UserID:        uint(userID),
//		LikesCount:    existingBlogPost.LikesCount,
//		Comments:      existingBlogPost.Comments,
//		CommentsCount: existingBlogPost.CommentsCount,
//	}
//	if blog.Title == "" {
//		blog.Title = existingBlogPost.Title
//	}
//	if blog.ContentText == "" {
//		blog.ContentText = existingBlogPost.ContentText
//	}
//	if blog.PhotoURL == "" {
//		blog.PhotoURL = existingBlogPost.PhotoURL
//	}
//	if blog.Description == "" {
//		blog.Description = existingBlogPost.Description
//	}
//	if blog.Category == "" {
//		blog.Category = existingBlogPost.Category
//	}
//	if err := ctr.svc.UpdateBlogPost(blog); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	return c.JSON(http.StatusOK, "Blog post updated successfully")
//}
//
//// DeleteBlogPost implements domain.BlogController.
//func (ctr *blogController) DeleteBlogPost(c echo.Context) error {
//	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	id, err := strconv.ParseUint(c.Param("postID"), 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Blog post not found")
//	}
//	if existingBlogPost.UserID != uint(userID) {
//		return c.JSON(http.StatusUnauthorized, "You are not authorized to delete this blog post")
//	}
//	if err := ctr.svc.DeleteBlogPost(uint(id)); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	return c.JSON(http.StatusOK, "Blog post deleted successfully")
//}
//
//// AddAndRemoveLike implements domain.BlogController.
//func (ctr *blogController) AddAndRemoveLike(c echo.Context) error {
//	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	tempID := c.Param("postID")
//	id, err := strconv.ParseUint(tempID, 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if existingBlogPost.ID == 0 {
//		return c.JSON(http.StatusNotFound, "Blog post not found")
//	}
//	s, err := ctr.svc.AddAndRemoveLike(&existingBlogPost, uint(userID))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if s == "like" {
//		return c.JSON(http.StatusOK, "Like Adeed to the Post")
//	}
//	return c.JSON(http.StatusOK, "Like Removed from the Post")
//}
//
//// AddComment implements domain.BlogController.
//func (ctr *blogController) AddComment(c echo.Context) error {
//	reqComment := &types.Comment{}
//	if err := c.Bind(reqComment); err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	if err := reqComment.Validate(); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	tempID := c.Param("postID")
//	id, err := strconv.ParseUint(tempID, 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if existingBlogPost.ID == 0 {
//		return c.JSON(http.StatusNotFound, "Blog post not found")
//	}
//	comment := &models.Comment{
//		Content:    reqComment.Content,
//		UserID:     uint(userID),
//		BlogPostID: existingBlogPost.ID,
//	}
//	if err := ctr.svc.AddComment(&existingBlogPost, comment); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	return c.JSON(http.StatusCreated, comment)
//}
//
//// GetCommentByUserID implements domain.BlogController.
//func (ctr *blogController) GetCommentByUserID(c echo.Context) error {
//	tempID := c.Param("postID")
//	id, err := strconv.ParseUint(tempID, 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	commentID, err := strconv.ParseUint(c.Param("commentID"), 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if existingBlogPost.ID == 0 {
//		return c.JSON(http.StatusNotFound, "Blog post not found")
//	}
//	comment, err := ctr.svc.GetCommentByUserID(&existingBlogPost, uint(commentID))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	return c.JSON(http.StatusOK, comment)
//}
//
//// GetComments implements domain.BlogController.
//func (ctr *blogController) GetComments(c echo.Context) error {
//	tempID := c.Param("postID")
//	id, err := strconv.ParseUint(tempID, 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if existingBlogPost.ID == 0 {
//		return c.JSON(http.StatusNotFound, "Blog post not found")
//	}
//	comments, err := ctr.svc.GetComments(&existingBlogPost)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	return c.JSON(http.StatusOK, comments)
//}
//
//// DeleteComment implements domain.BlogController.
//func (ctr *blogController) DeleteComment(c echo.Context) error {
//	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	tempID := c.Param("postID")
//	id, err := strconv.ParseUint(tempID, 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	commentID, err := strconv.ParseUint(c.Param("commentID"), 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if existingBlogPost.ID == 0 {
//		return c.JSON(http.StatusNotFound, "Blog post not found")
//	}
//	if _, err := ctr.svc.GetCommentByUserID(&existingBlogPost, uint(userID)); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if err := ctr.svc.DeleteComment(&existingBlogPost, uint(commentID)); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	return c.JSON(http.StatusOK, "Comment deleted successfully")
//}
//
//// UpdateComment implements domain.BlogController.
//func (ctr *blogController) UpdateComment(c echo.Context) error {
//	reqComment := &types.Comment{}
//	if err := c.Bind(reqComment); err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	if err := reqComment.Validate(); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	tempID := c.Param("postID")
//	id, err := strconv.ParseUint(tempID, 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	commentID, err := strconv.ParseUint(c.Param("commentID"), 10, 64)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, "Invalid data request")
//	}
//	existingBlogPost, err := ctr.svc.GetBlogPost(uint(id))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if existingBlogPost.ID == 0 {
//		return c.JSON(http.StatusNotFound, "Blog post not found")
//	}
//	if _, err := ctr.svc.GetCommentByUserID(&existingBlogPost, uint(userID)); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	comment := &models.Comment{
//		Model:   gorm.Model{ID: uint(commentID)},
//		Content: reqComment.Content,
//	}
//	if err := ctr.svc.UpdateComment(&existingBlogPost, comment); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	return c.JSON(http.StatusOK, "Comment updated successfully")
//}
