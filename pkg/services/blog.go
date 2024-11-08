package services

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/types"
	blogconsts "Blog_API/pkg/utils/consts/blog"
	userconsts "Blog_API/pkg/utils/consts/user"
	"errors"
	"github.com/google/uuid"
	"time"
)

// Parent struct to implement interface binding
type blogService struct {
	repo domain.BlogRepository
	uSvc domain.Service
}

// Interface binding
func NewBlogService(repo domain.BlogRepository, usvc domain.Service) domain.BlogService {
	return &blogService{
		repo: repo,
		uSvc: usvc,
	}
}

// CreateBlogPost implements domain.BlogService.
func (svc *blogService) CreateBlogPost(reqBlogPost types.BlogPostRequest, userID string) (types.BlogResp, error) {

	user, err := svc.uSvc.GetUser(userID)
	if err != nil {
		return types.BlogResp{}, err
	}

	if user.ID == "" {
		return types.BlogResp{}, errors.New(userconsts.ErrorGettingUser)
	}

	reqBlog := models.BlogPost{
		ID:          uuid.NewString(),
		UserID:      user.ID,
		Title:       reqBlogPost.Title,
		ContentText: reqBlogPost.ContentText,
		PhotoURL:    reqBlogPost.PhotoURL,
		Description: reqBlogPost.Description,
		Category:    reqBlogPost.Category,
		IsPublished: reqBlogPost.IsPublished,
		PublishedAt: time.Now(),
	}

	if createBlogErr := svc.repo.CreateBlogPost(reqBlog); createBlogErr != nil {
		return types.BlogResp{}, createBlogErr
	}

	return convertBlogPostToBlogResp(reqBlog), nil
}

// GetBlogPost implements domain.BlogService.
func (svc *blogService) GetBlogPost(blogID string) (types.BlogResp, error) {
	blogPost, err := svc.repo.GetBlogPost(blogID)
	if err != nil {
		return types.BlogResp{}, err
	}

	if blogPost.ID == "" {
		return types.BlogResp{}, errors.New(userconsts.ErrorGettingUser)
	}

	return convertBlogPostToBlogResp(blogPost), nil
}

// GetBlogPosts implements domain.BlogService.
func (svc *blogService) GetBlogPosts() ([]types.BlogResp, error) {

	var blogResp []types.BlogResp
	blogPosts, err := svc.repo.GetBlogPosts()
	if err != nil {
		return blogResp, err
	}

	for _, blogPost := range blogPosts {
		blogResp = append(blogResp, convertBlogPostToBlogResp(blogPost))
	}

	return blogResp, nil
}

// GetBlogPosts implements domain.BlogService.
func (svc *blogService) GetBlogPostsOfUser(userID string, blogIDs []string) ([]types.BlogResp, error) {

	var blogResp []types.BlogResp
	blogPosts, err := svc.repo.GetBlogPostsOfUser(userID, blogIDs)
	if err != nil {
		return []types.BlogResp{}, err
	}

	for _, blogPost := range blogPosts {
		blogResp = append(blogResp, convertBlogPostToBlogResp(blogPost))
	}

	return blogResp, nil
}

// UpdateBlogPost implements domain.BlogService.
func (svc *blogService) UpdateBlogPost(userID string, blogID string, blogPostReq types.UpdateBlogPostRequest) (types.BlogResp, error) {

	user, err := svc.uSvc.GetUser(userID)
	if err != nil {
		return types.BlogResp{}, err
	}

	blogPost, err := svc.repo.GetBlogPostsOfUser(userID, []string{blogID})
	if err != nil {
		return types.BlogResp{}, err
	}

	if len(blogPost) == 0 {
		return types.BlogResp{}, errors.New(blogconsts.ErrorGettingBlog)
	}

	blog := models.BlogPost{
		ID:          blogPost[0].ID,
		UserID:      user.ID,
		Title:       blogPostReq.Title,
		ContentText: blogPostReq.ContentText,
		PhotoURL:    blogPostReq.PhotoURL,
		Description: blogPostReq.Description,
		Category:    blogPostReq.Category,
		IsPublished: blogPostReq.IsPublished,
		PublishedAt: time.Now(),
	}

	if updateErr := svc.repo.UpdateBlogPost(blog); updateErr != nil {
		return types.BlogResp{}, updateErr
	}

	return convertBlogPostToBlogResp(blog), nil
}

// DeleteBlogPost implements domain.BlogService.
func (svc *blogService) DeleteBlogPost(userID string, blogID string) error {

	user, err := svc.uSvc.GetUser(userID)
	if err != nil {
		return err
	}

	blogPost, err := svc.repo.GetBlogPostsOfUser(user.ID, []string{blogID})
	if err != nil {
		return err
	}

	if len(blogPost) == 0 {
		return errors.New(blogconsts.YouAreNotAuthorizedToDeleteThisBlog)
	}

	if deleteErr := svc.repo.DeleteBlogPost(blogID); deleteErr != nil {
		return deleteErr
	}

	return nil
}

// AddAndRemoveReaction implements domain.BlogService.
func (svc *blogService) AddAndRemoveReaction(userID string, blogID string, reactionID uint64) (types.BlogResp, error) {

	if blogconsts.ReactionTypes[reactionID] == "" {
		return types.BlogResp{}, errors.New(blogconsts.InvalidReactionID)
	}

	user, err := svc.uSvc.GetUser(userID)
	if err != nil {
		return types.BlogResp{}, err
	}

	blogPost, err := svc.repo.GetBlogPost(blogID)
	if err != nil {
		return types.BlogResp{}, err
	}

	if blogPost.ID == "" {
		return types.BlogResp{}, errors.New(blogconsts.ErrorGettingBlog)
	}

	blogPost, err = svc.repo.AddAndRemoveReaction(user.ID, reactionID, blogPost)
	if err != nil {
		return types.BlogResp{}, err
	}

	return convertBlogPostToBlogResp(blogPost), nil
}

//// AddComment implements domain.BlogService.
//func (svc *blogService) AddComment(blogPost *models.BlogPost, comment *models.Comment) error {
//	if err := svc.repo.AddCommentRepo(blogPost, comment); err != nil {
//		return err
//	}
//	return nil
//}
//
//// GetCommentByUserID implements domain.BlogService.
//func (svc *blogService) GetCommentByUserID(blogPost *models.BlogPost, commentID uint) (models.Comment, error) {
//	comment, err := svc.repo.GetCommentByUserIDRepo(blogPost, commentID)
//	if err != nil {
//		return comment, err
//	}
//	return comment, nil
//}
//
//// GetComments implements domain.BlogService.
//func (svc *blogService) GetComments(blogPost *models.BlogPost) ([]models.Comment, error) {
//	comments, err := svc.repo.GetCommentsRepo(blogPost)
//	if err != nil {
//		return comments, err
//	}
//	return comments, nil
//}
//
//// DeleteComment implements domain.BlogService.
//func (svc *blogService) DeleteComment(blogPost *models.BlogPost, commentID uint) error {
//	if err := svc.repo.DeleteCommentRepo(blogPost, commentID); err != nil {
//		return err
//	}
//	return nil
//}
//
//// UpdateComment implements domain.BlogService.
//func (svc *blogService) UpdateComment(blogPost *models.BlogPost, comment *models.Comment) error {
//	if err := svc.repo.UpdateCommentRepo(blogPost, comment); err != nil {
//		return err
//	}
//	return nil
//}

func convertBlogPostToBlogResp(blogPost models.BlogPost) types.BlogResp {
	return types.BlogResp{
		ID:             blogPost.ID,
		UserID:         blogPost.UserID,
		Title:          blogPost.Title,
		ContentText:    blogPost.ContentText,
		PhotoURL:       blogPost.PhotoURL,
		Description:    blogPost.Description,
		Category:       blogPost.Category,
		CommentsCount:  blogPost.CommentsCount,
		Comments:       blogPost.Comments,
		ReactionsCount: blogPost.ReactionsCount,
		Reactions:      blogPost.Reactions,
		Views:          blogPost.Views,
		IsPublished:    blogPost.IsPublished,
		PublishedAt:    blogPost.PublishedAt.Format(time.RFC3339),
	}
}
