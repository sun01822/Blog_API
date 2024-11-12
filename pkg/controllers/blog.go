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
	"net/http"
	"strconv"
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
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Param blogPost body types.BlogPostRequest true "Blog Post Request"
// @Success 200 {object} types.BlogResp "blog post created successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error creating blog"
// @Router /blog/create [post]
func (ctr *blogController) CreateBlogPost(c echo.Context) error {

	userID, err := extractUserID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if userID == "" {
		return response.ErrorResponse(c, errors.New(userconsts.UserIDRequired), consts.InvalidDataRequest)
	}

	reqBlogPost := types.BlogPostRequest{}
	if bindErr := c.Bind(&reqBlogPost); bindErr != nil {
		return response.ErrorResponse(c, bindErr, consts.InvalidDataRequest)
	}

	if validationErr := reqBlogPost.Validate(); validationErr != nil {
		return response.ErrorResponse(c, validationErr, consts.ValidationError)
	}

	blog, err := ctr.svc.CreateBlogPost(reqBlogPost, userID)
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
// @Success 200 {object} types.BlogResp "blog fetched successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error getting blog"
// @Router /blog/get [get]
func (ctr *blogController) GetBlogPost(c echo.Context) error {

	reqBlogID, err := extractBlogID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

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
// @Success 200 {array} types.BlogResp "blogs fetched successfully"
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
// @Param Authorization header string true "Bearer <token>"
// @Param blog_id query string true "Blog ID"
// @Param user_id query string true "User ID"
// @Param blog_ids query string false "Blog IDs"
// @Success 200 {array} types.BlogResp "blogs fetched successfully"
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
// @Success 200 {object} types.BlogResp "blog updated successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error updating blog"
// @Router /blog/update [put]
func (ctr *blogController) UpdateBlogPost(c echo.Context) error {

	userID, reqBlogID, err := extractUserIDAndReqBlogID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := checkUserIDAndBlogIDIsEmptyOrNot(userID, reqBlogID); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	updateBlogReq := types.UpdateBlogPostRequest{}
	if err := c.Bind(&updateBlogReq); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := updateBlogReq.Validate(); err != nil {
		return response.ErrorResponse(c, err, consts.ValidationError)
	}

	blog, err := ctr.svc.UpdateBlogPost(userID, reqBlogID, updateBlogReq)
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorUpdatingBlog)
	}

	return response.SuccessResponse(c, blogconsts.BlogUpdatedSuccessfully, blog)
}

// DeleteBlogPost implements domain.BlogController.
// @Summary Delete a blog post
// @Description Delete a blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Param blog_id query string true "Blog ID"
// @Success 200 {string} string "blog deleted successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error deleting blog"
// @Router /blog/delete [delete]
func (ctr *blogController) DeleteBlogPost(c echo.Context) error {

	userID, reqBlogID, err := extractUserIDAndReqBlogID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := checkUserIDAndBlogIDIsEmptyOrNot(userID, reqBlogID); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := ctr.svc.DeleteBlogPost(userID, reqBlogID); err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorDeletingBlog)
	}

	return response.SuccessResponse(c, blogconsts.BlogDeletedSuccessfully, nil)
}

// AddAndRemoveReaction implements domain.BlogController.
// @Summary Add or remove a reaction
// @Description Add or remove a reaction
// @Tags Blog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Param blog_id query string true "Blog ID"
// @Param reaction_id query uint64 true "Reaction ID"
// @Success 200 {object} types.BlogResp "reaction added successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error adding or removing reaction"
// @Router /blog/reaction [post]
func (ctr *blogController) AddAndRemoveReaction(c echo.Context) error {

	userID, reqBlogID, err := extractUserIDAndReqBlogID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := checkUserIDAndBlogIDIsEmptyOrNot(userID, reqBlogID); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	reqReactionID, err := extractReqReactionID(c)

	resp, err := ctr.svc.AddAndRemoveReaction(userID, reqBlogID, reqReactionID)
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorAddingRemovingReaction)
	}

	return response.SuccessResponse(c, blogconsts.ReactionAddedSuccessfully, resp)
}

// AddComment implements domain.BlogController.
// @Summary Add a comment
// @Description Add a comment
// @Tags Blog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Param blog_id query string true "Blog ID"
// @Param comment body types.Comment true "Comment"
// @Success 200 {object} types.BlogResp "comment added successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error adding comment"
// @Router /blog/comment [post]
func (ctr *blogController) AddComment(c echo.Context) error {

	userID, reqBlogID, err := extractUserIDAndReqBlogID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := checkUserIDAndBlogIDIsEmptyOrNot(userID, reqBlogID); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	reqComment := types.Comment{}
	if bindErr := c.Bind(&reqComment); bindErr != nil {
		return response.ErrorResponse(c, bindErr, consts.InvalidDataRequest)
	}

	if validationErr := reqComment.Validate(); validationErr != nil {
		return response.ErrorResponse(c, validationErr, consts.ValidationError)
	}

	resp, err := ctr.svc.AddComment(userID, reqBlogID, reqComment)
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorAddingComment)
	}

	return response.SuccessResponse(c, blogconsts.CommentAddedSuccessfully, resp)
}

// GetComments implements domain.BlogController.
// @Summary Get comments of a blog post
// @Description Get comments of a blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Param blog_id query string true "Blog ID"
// @Param comment_ids query string false "Comment IDs"
// @Success 200 {array} types.CommentResp "comments fetched successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error getting comments"
// @Router /blog/comments [get]
func (ctr *blogController) GetComments(c echo.Context) error {

	userID, reqBlogID, err := extractUserIDAndReqBlogID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := checkUserIDAndBlogIDIsEmptyOrNot(userID, reqBlogID); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	reqCommentIDs := extractReqCommentIDs(c)

	comments, err := ctr.svc.GetComments(userID, reqBlogID, reqCommentIDs)
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorGettingComments)
	}

	return response.SuccessResponse(c, blogconsts.CommentsFetchSuccessfully, comments)
}

// DeleteComment implements domain.BlogController.
// @Summary Delete a comment
// @Description Delete a comment
// @Tags Blog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Param blog_id query string true "Blog ID"
// @Param comment_id query string true "Comment ID"
// @Success 200 {string} string "comment deleted successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error deleting comment"
// @Router /blog/comment [delete]
func (ctr *blogController) DeleteComment(c echo.Context) error {

	userID, reqBlogID, err := extractUserIDAndReqBlogID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := checkUserIDAndBlogIDIsEmptyOrNot(userID, reqBlogID); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	reqCommentID, err := extractReqCommentID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := checkReqCommentIDIsEmptyOrNot(reqCommentID); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := ctr.svc.DeleteComment(userID, reqBlogID, reqCommentID); err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorDeletingComment)
	}

	return response.SuccessResponse(c, blogconsts.CommentDeletedSuccessfully, nil)
}

// UpdateComment implements domain.BlogController.
// @Summary Update a comment
// @Description Update a comment
// @Tags Blog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer <token>"
// @Param blog_id query string true "Blog ID"
// @Param comment_id query string true "Comment ID"
// @Param comment body types.Comment true "Comment"
// @Success 200 {object} types.BlogResp "comment updated successfully"
// @Failure 400 {string} string "invalid data request"
// @Failure 500 {string} string "error updating comment"
// @Router /blog/comment [put]
func (ctr *blogController) UpdateComment(c echo.Context) error {

	userID, reqBlogID, err := extractUserIDAndReqBlogID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := checkUserIDAndBlogIDIsEmptyOrNot(userID, reqBlogID); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	reqCommentID, err := extractReqCommentID(c)
	if err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	if err := checkReqCommentIDIsEmptyOrNot(reqCommentID); err != nil {
		return response.ErrorResponse(c, err, consts.InvalidDataRequest)
	}

	reqComment := types.Comment{}
	if err := c.Bind(&reqComment); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data request")
	}

	if err := reqComment.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := ctr.svc.UpdateComment(userID, reqBlogID, reqCommentID, reqComment)
	if err != nil {
		return response.ErrorResponse(c, err, blogconsts.ErrorUpdatingComment)
	}

	return response.SuccessResponse(c, blogconsts.CommentUpdatedSuccessfully, resp)
}

func extractUserIDAndBlogIDs(ctx echo.Context) (string, []string, error) {

	userID, parseErr := uuid.Parse(ctx.Get(userconsts.UserID).(string))
	if parseErr != nil {
		return "", nil, response.ErrorResponse(ctx, parseErr, consts.InvalidDataRequest)
	}

	blogIDsParam := ctx.QueryParam(blogconsts.BlogIDs)
	blogIDs := strings.Fields(strings.ReplaceAll(blogIDsParam, ",", " "))

	return userID.String(), blogIDs, nil
}

func extractUserIDAndReqBlogID(ctx echo.Context) (string, string, error) {

	userID, parseErr := uuid.Parse(ctx.Get(userconsts.UserID).(string))
	if parseErr != nil {
		return "", "", response.ErrorResponse(ctx, parseErr, consts.InvalidDataRequest)
	}

	reqBlogID, parseErr := uuid.Parse(ctx.QueryParam(blogconsts.BlogID))
	if parseErr != nil {
		return "", "", response.ErrorResponse(ctx, parseErr, consts.InvalidDataRequest)
	}

	return userID.String(), reqBlogID.String(), nil
}

func extractUserID(ctx echo.Context) (string, error) {

	userID, parseErr := uuid.Parse(ctx.Get(userconsts.UserID).(string))
	if parseErr != nil {
		return "", response.ErrorResponse(ctx, parseErr, consts.InvalidDataRequest)
	}

	return userID.String(), nil
}

func extractBlogID(ctx echo.Context) (string, error) {

	reqBlogID, parseErr := uuid.Parse(ctx.QueryParam(blogconsts.BlogID))
	if parseErr != nil {
		return "", response.ErrorResponse(ctx, parseErr, consts.InvalidDataRequest)
	}

	return reqBlogID.String(), nil
}

func checkUserIDAndBlogIDIsEmptyOrNot(userID, reqBlogID string) error {

	if userID == "" {
		return errors.New(userconsts.UserIDRequired)
	}

	if reqBlogID == "" {
		return errors.New(blogconsts.BlogIDRequired)
	}

	return nil
}

func extractReqCommentID(ctx echo.Context) (string, error) {

	reqCommentID, parseErr := uuid.Parse(ctx.QueryParam(blogconsts.CommentID))
	if parseErr != nil {
		return "", response.ErrorResponse(ctx, parseErr, consts.InvalidDataRequest)
	}

	return reqCommentID.String(), nil
}

func checkReqCommentIDIsEmptyOrNot(reqCommentID string) error {

	if reqCommentID == "" {
		return errors.New(blogconsts.InvalidCommentID)
	}

	return nil
}

func extractReqReactionID(ctx echo.Context) (uint64, error) {

	reqReactionID, parseErr := strconv.ParseUint(ctx.QueryParam(blogconsts.ReactionID), 10, 64)
	if parseErr != nil {
		return 0, response.ErrorResponse(ctx, parseErr, consts.InvalidDataRequest)
	}

	return reqReactionID, nil
}

func extractReqCommentIDs(ctx echo.Context) []string {

	reqCommentIDsParam := ctx.QueryParam(blogconsts.CommentIDs)
	reqCommentIDs := strings.Fields(strings.ReplaceAll(reqCommentIDsParam, ",", " "))

	return reqCommentIDs
}
