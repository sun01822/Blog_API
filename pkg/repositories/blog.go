package repositories

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
	"Blog_API/pkg/utils/consts"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Parent struct to implement interface binding
type blogRepo struct {
	d *gorm.DB
}

// Interface binding
func NewBlogRepo(db *gorm.DB) domain.BlogRepository {
	return &blogRepo{
		d: db,
	}
}

// CreateBlogPost implements domain.BlogRepository.
func (repo *blogRepo) CreateBlogPost(blogPost models.BlogPost) error {
	err := repo.d.Preload(consts.REACTIONS).Preload(consts.COMMENTS).Create(&blogPost).Error
	if err != nil {
		return err
	}
	return nil
}

// GetBlogPost implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPost(blogID string) (models.BlogPost, error) {

	var blogPost models.BlogPost
	err := repo.d.Preload(consts.REACTIONS).Preload(consts.COMMENTS).Where("id = ?", blogID).First(&blogPost).Error
	if err != nil {
		return blogPost, err
	}

	return blogPost, nil
}

// GetBlogPosts implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPosts() ([]models.BlogPost, error) {

	var blogPosts []models.BlogPost
	err := repo.d.Preload(consts.REACTIONS).Preload(consts.COMMENTS).Find(&blogPosts).Error
	if err != nil {
		return blogPosts, err
	}

	return blogPosts, nil
}

// GetBlogPostsBasedOnCategory implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPostsBasedOnCategory(category string) ([]models.BlogPost, error) {

	var blogPosts []models.BlogPost
	err := repo.d.Preload(consts.REACTIONS).Preload(consts.COMMENTS).Where("category = ?", category).Find(&blogPosts).Error
	if err != nil {
		return blogPosts, err
	}

	return blogPosts, nil
}

// GetBlogPosts implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPostsOfUser(userID string, blogIDs []string) ([]models.BlogPost, error) {

	var blogPosts []models.BlogPost
	query := repo.d.Preload(consts.REACTIONS).Preload(consts.COMMENTS).Where("user_id = ?", userID)

	if len(blogIDs) > 0 {
		query = query.Where("id IN ?", blogIDs)
	}

	err := query.Find(&blogPosts).Error
	if err != nil {
		return blogPosts, err
	}

	return blogPosts, nil
}

// UpdateBlogPost implements domain.BlogRepository.
func (repo *blogRepo) UpdateBlogPost(blogPost models.BlogPost) error {

	err := repo.d.Preload(consts.REACTIONS).Preload(consts.COMMENTS).Updates(&blogPost).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteBlogPost implements domain.BlogRepository.
func (repo *blogRepo) DeleteBlogPost(blogID string) error {

	err := repo.d.Preload(consts.REACTIONS).Preload(consts.COMMENTS).Where("id = ?", blogID).Delete(&models.BlogPost{}).Error
	if err != nil {
		return err
	}

	return nil
}

// AddAndRemoveReaction implements domain.BlogRepository.
func (repo *blogRepo) AddAndRemoveReaction(userID string, reactionID uint64, blogPost models.BlogPost) (models.BlogPost, error) {

	tx, err := beginTransaction(repo.d)
	if err != nil {
		return models.BlogPost{}, err
	}

	reaction, err := repo.findReaction(tx, userID, blogPost.ID)
	if err != nil {
		if createErr := repo.createReaction(tx, userID, reactionID, &blogPost); createErr != nil {
			return models.BlogPost{}, createErr
		}
	} else {
		if reaction.Type == reactionID {
			if removeErr := repo.removeReaction(tx, &reaction, &blogPost); removeErr != nil {
				return models.BlogPost{}, removeErr
			}
		} else {
			if updateErr := repo.updateReaction(tx, reaction, reactionID, &blogPost); updateErr != nil {
				return models.BlogPost{}, updateErr
			}
		}
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return models.BlogPost{}, commitErr
	}

	return blogPost, nil
}

// AddComment implements domain.BlogRepository.
func (repo *blogRepo) AddComment(blogPost models.BlogPost, comment models.Comment) (models.BlogPost, error) {

	tx, err := beginTransaction(repo.d)
	if err != nil {
		return models.BlogPost{}, err
	}

	commentErr := tx.Create(&comment).Error
	if commentErr != nil {
		return models.BlogPost{}, commentErr
	}

	if appendErr := tx.Model(&blogPost).Association(consts.COMMENTS).Append(&comment); appendErr != nil {
		return models.BlogPost{}, appendErr
	}

	if updateCountErr := tx.Model(&blogPost).Update(consts.CommentCounts, blogPost.CommentsCount+1).Error; updateCountErr != nil {
		return models.BlogPost{}, updateCountErr
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return models.BlogPost{}, commitErr
	}

	return blogPost, nil
}

// GetComments implements domain.BlogRepository.
func (repo *blogRepo) GetComments(blogID string, commentIDs []string) ([]models.Comment, error) {

	var comments []models.Comment
	query := repo.d.Where("blog_post_id = ?", blogID)

	if len(commentIDs) != 0 {
		query = query.Where("id in (?)", commentIDs)
	}

	err := query.Find(&comments).Error
	if err != nil {
		return comments, err
	}

	return comments, nil
}

// DeleteComment implements domain.BlogRepository.
func (repo *blogRepo) DeleteComment(blogPost models.BlogPost, commentID string) error {

	tx, err := beginTransaction(repo.d)
	if err != nil {
		return err
	}

	var comment models.Comment
	err = tx.Where("id = ? AND blog_post_id = ?", commentID, blogPost.ID).First(&comment).Error
	if err != nil {
		return err
	}

	err = tx.Delete(&comment).Error
	if err != nil {
		return err
	}

	err = tx.Model(&blogPost).Association(consts.COMMENTS).Delete(&comment)
	if err != nil {
		return err
	}

	err = tx.Model(&blogPost).Update(consts.CommentCounts, blogPost.CommentsCount-1).Error
	if err != nil {
		return err
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return commitErr
	}

	return nil
}

// UpdateComment implements domain.BlogRepository.
func (repo *blogRepo) UpdateComment(blogPost models.BlogPost, comment models.Comment) (models.BlogPost, error) {

	tx, err := beginTransaction(repo.d)
	if err != nil {
		return models.BlogPost{}, err
	}

	err = tx.Updates(&comment).Error
	if err != nil {
		return models.BlogPost{}, err
	}

	updatedComment, err := repo.GetComments(blogPost.ID, []string{comment.ID})

	// Update the blog post with the updated comment
	err = tx.Model(&blogPost).Association(consts.COMMENTS).Replace(&updatedComment)
	if err != nil {
		return models.BlogPost{}, err
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		return models.BlogPost{}, commitErr
	}

	return blogPost, nil
}

func beginTransaction(db *gorm.DB) (*gorm.DB, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil || tx.Error != nil {
			tx.Rollback()
		}
	}()

	return tx, nil
}

func (repo *blogRepo) findReaction(tx *gorm.DB, userID, blogPostID string) (models.Reaction, error) {

	var reaction models.Reaction
	err := tx.Where("blog_post_id = ? AND user_id = ?", blogPostID, userID).First(&reaction).Error
	if err != nil {
		return reaction, err
	}

	return reaction, err
}

func (repo *blogRepo) createReaction(tx *gorm.DB, userID string, reactionID uint64, blogPost *models.BlogPost) error {

	reaction := models.Reaction{
		ID:         uuid.NewString(),
		UserID:     userID,
		BlogPostID: blogPost.ID,
		Type:       reactionID,
	}

	if err := tx.Create(&reaction).Error; err != nil {
		return err
	}

	if err := tx.Model(&blogPost).Association(consts.REACTIONS).Append(&reaction); err != nil {
		return err
	}

	if err := tx.Model(&blogPost).Update(consts.ReactionCounts, blogPost.ReactionsCount+1).Error; err != nil {
		return err
	}

	return nil
}

func (repo *blogRepo) removeReaction(tx *gorm.DB, reaction *models.Reaction, blogPost *models.BlogPost) error {

	if err := tx.Delete(&reaction).Where("id = ?", reaction.ID).Error; err != nil {
		return err
	}

	if err := tx.Model(&blogPost).Association(consts.REACTIONS).Delete(&reaction); err != nil {
		return err
	}

	if err := tx.Model(&blogPost).Update(consts.ReactionCounts, blogPost.ReactionsCount-1).Error; err != nil {
		return err
	}

	return nil
}

func (repo *blogRepo) updateReaction(tx *gorm.DB, reaction models.Reaction, reactionID uint64, blogPost *models.BlogPost) error {

	if err := tx.Model(&reaction).Update("type", reactionID).Error; err != nil {
		return err
	}

	if err := tx.Model(&blogPost).Association(consts.REACTIONS).Replace(&reaction); err != nil {
		return err
	}

	return nil
}
