package services

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/types"
	userconsts "Blog_API/pkg/utils/consts/user"
	"errors"
	"github.com/google/uuid"
	"time"
)

// Parent struct to implement interface binding
type blogService struct {
	repo  domain.BlogRepository
	uRepo domain.Repository
}

// Interface binding
func NewBlogService(repo domain.BlogRepository, uRepo domain.Repository) domain.BlogService {
	return &blogService{
		repo:  repo,
		uRepo: uRepo,
	}
}

// CreateBlogPost implements domain.BlogService.
func (svc *blogService) CreateBlogPost(reqBlogPost types.BlogPostRequest, userID string) (types.BlogResp, error) {

	user, err := svc.uRepo.GetUser(userID)
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

//// GetBlogPost implements domain.BlogService.
//func (svc *blogService) GetBlogPost(id uint) (models.BlogPost, error) {
//	blogPost, err := svc.repo.GetBlogPostRepo(id)
//	if err != nil {
//		return blogPost, err
//	}
//	return blogPost, nil
//}
//
//// GetBlogPosts implements domain.BlogService.
//func (svc *blogService) GetBlogPosts() ([]models.BlogPost, error) {
//	blogPosts, err := svc.repo.GetBlogPostsRepo()
//	if err != nil {
//		return blogPosts, err
//	}
//	return blogPosts, nil
//}
//
//// GetBlogPosts implements domain.BlogService.
//func (svc *blogService) GetBlogPostsOfUser(userID uint) ([]models.BlogPost, error) {
//	blogPosts, err := svc.repo.GetBlogPostsOfUserRepo(userID)
//	if err != nil {
//		return blogPosts, err
//	}
//	return blogPosts, nil
//}
//
//// UpdateBlogPost implements domain.BlogService.
//func (svc *blogService) UpdateBlogPost(blogPost *models.BlogPost) error {
//	if err := svc.repo.UpdateBlogPostRepo(blogPost); err != nil {
//		return err
//	}
//	return nil
//}
//
//// DeleteBlogPost implements domain.BlogService.
//func (svc *blogService) DeleteBlogPost(id uint) error {
//	if err := svc.repo.DeleteBlogPostRepo(id); err != nil {
//		return err
//	}
//	return nil
//}
//
//// AddAndRemoveLike implements domain.BlogService.
//func (svc *blogService) AddAndRemoveLike(blogPost *models.BlogPost, userID uint) (string, error) {
//	s, err := svc.repo.AddAndRemoveLikeRepo(blogPost, userID)
//	if err != nil {
//		return s, err
//	}
//	return s, nil
//}
//
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
		UserID:         blogPost.UserID,
		Title:          blogPost.Title,
		ContentText:    blogPost.ContentText,
		PhotoURL:       blogPost.PhotoURL,
		Description:    blogPost.Description,
		Category:       blogPost.Category,
		CommentsCount:  blogPost.CommentsCount,
		ReactionsCount: blogPost.ReactionsCount,
		Views:          blogPost.Views,
		IsPublished:    blogPost.IsPublished,
		PublishedAt:    blogPost.PublishedAt.Format(time.RFC3339),
	}
}
