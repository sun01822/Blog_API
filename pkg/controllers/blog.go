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
	"strings"
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
// @Router /blog/create [post]
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
// @Summary Get a blog post
// @Description Get a blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Param blog_id query string true "Blog ID"
// @Success 200 {object} types.BlogResp "Blog Fetched Successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error getting blog"
// @Router /blog/get [get]
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
// @Router /blog/getAll [get]
func (ctr *blogController) GetBlogPosts(c echo.Context) error {

	blogPosts, err := ctr.svc.GetBlogPosts()
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorGettingBlogs)
	}

	return response.SuccessResponse(c, blogconsts.BlogsFetchSuccessfully, blogPosts)
}

// GetBlogPostsOfUser implements domain.BlogController.
// @Summary Get all blog posts of a user
// @Description Get all blog posts of a user
// @Tags Blog
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Param blog_ids query string false "Blog IDs"
// @Success 200 {array} types.BlogResp "Blogs Fetched Successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error getting blogs"
// @Router /blog/get/user [get]
func (ctr *blogController) GetBlogPostsOfUser(c echo.Context) error {

	userID, blogIDs, err := extractUserIDAndBlogIDs(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	blogPosts, err := ctr.svc.GetBlogPostsOfUser(userID, blogIDs)
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorGettingBlogs)
	}

	return response.SuccessResponse(c, blogconsts.BlogsFetchSuccessfullyOfUser, blogPosts)
}

// UpdateBlogPost implements domain.BlogController.
// @Summary Update a blog post
// @Description Update a blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Param blog_id query string true "Blog ID"
// @Param blogPost body types.UpdateBlogPostRequest true "update blog post request"
// @Success 200 {object} types.BlogResp "blog Updated Successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error updating blog"
// @Router /blog/update [put]
func (ctr *blogController) UpdateBlogPost(c echo.Context) error {

	userID, parseErr := uuid.Parse(c.Get(userconsts.UserID).(string))
	if parseErr != nil {
		return response.ErrorResponse(c, parseErr, consts.InvalidDataRequest)
	}

	if userID.String() == "" {
		return response.ErrorResponse(c, errors.New(userconsts.UserIDRequired), consts.InvalidDataRequest)
	}

	blogID := c.QueryParam(blogconsts.BlogID)
	if blogID == "" {
		return response.ErrorResponse(c, errors.New(blogconsts.BlogIDRequired), consts.InvalidDataRequest)
	}

	updateBlogReq := types.UpdateBlogPostRequest{}
	if err := c.Bind(&updateBlogReq); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := updateBlogReq.Validate(); err != nil {
		return response.ErrorResponse(c, err, consts.ValidationError)
	}

	blog, err := ctr.svc.UpdateBlogPost(userID.String(), blogID, updateBlogReq)
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorUpdatingBlog)
	}

	return response.SuccessResponse(c, blogconsts.BlogUpdatedSuccessfully, blog)
}

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

func extractUserIDAndBlogIDs(ctx echo.Context) (string, []string, error) {

	userID := ctx.Get(userconsts.UserID).(string)
	_, err := uuid.Parse(userID)
	if err != nil {
		return "", nil, errors.New(consts.InvalidDataRequest)
	}

	blogIDsParam := ctx.QueryParam(blogconsts.BlogIDs)
	blogIDs := strings.Fields(strings.ReplaceAll(blogIDsParam, ",", " "))

	return userID, blogIDs, nil
}
