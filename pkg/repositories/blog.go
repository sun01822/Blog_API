package repositories

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/models"
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
func (repo *blogRepo) CreateBlogPostRepo(blogPost *models.BlogPost) error {
	err := repo.d.Preload("Likes").Preload("Comments").Create(blogPost).Error
	if err != nil {
		return err
	}
	return nil
}

// GetBlogPost implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPostRepo(id uint) (models.BlogPost, error) {
	var blogPost models.BlogPost
	err := repo.d.Preload("Likes").Preload("Comments").Where("id = ?", id).First(&blogPost).Error
	if err != nil {
		return blogPost, err
	}
	return blogPost, nil
}

// GetBlogPosts implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPostsRepo() ([]models.BlogPost, error) {
	var blogPosts []models.BlogPost
	err := repo.d.Preload("Likes").Preload("Comments").Find(&blogPosts).Error
	if err != nil {
		return blogPosts, err
	}
	return blogPosts, nil
}

// GetBlogPosts implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPostsOfUserRepo(userID uint) ([]models.BlogPost, error) {
	var blogPosts []models.BlogPost
	err := repo.d.Preload("Likes").Preload("Comments").Where("user_id = ?", userID).Find(&blogPosts).Error
	if err != nil {
		return blogPosts, err
	}
	return blogPosts, nil
}

// UpdateBlogPost implements domain.BlogRepository.
func (repo *blogRepo) UpdateBlogPostRepo(blogPost *models.BlogPost) error {
	err := repo.d.Preload("Likes").Preload("Comments").Save(blogPost).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteBlogPost implements domain.BlogRepository.
func (repo *blogRepo) DeleteBlogPostRepo(id uint) error {
	err := repo.d.Preload("Likes").Preload("Comments").Where("id = ?", id).Delete(&models.BlogPost{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *blogRepo) AddAndRemoveLikeRepo(blogPost *models.BlogPost, userID uint) (string, error) {
	// check if the user has already liked the post
	// if yes, remove the like
	// if no, add the like
	// return the updated blogPost
	var like models.Like
	err := repo.d.Where("user_id = ? AND blog_post_id = ?", userID, blogPost.ID).First(&like).Error
	if err != nil {
		// user has not liked the post
		like = models.Like{
			UserID:     userID,
			BlogPostID: blogPost.ID,
		}
		err = repo.d.Create(&like).Error
		if err != nil {
			return "", err
		}
		repo.d.Model(blogPost).Association("Likes").Append(&like)
		repo.d.Model(blogPost).Update("likes_count", blogPost.LikesCount+1)
		return "like", nil
	} else {
		// user has liked the post
		err = repo.d.Delete(&like).Error
		if err != nil {
			return "", err
		}
		repo.d.Model(blogPost).Association("Likes").Delete(&like)
		repo.d.Model(blogPost).Update("likes_count", blogPost.LikesCount-1)
		return "remove", nil
	}
}

// AddComment implements domain.BlogRepository.
func (repo *blogRepo) AddCommentRepo(blogPost *models.BlogPost, comment *models.Comment) error {
	err := repo.d.Create(comment).Error
	if err != nil {
		return err
	}
	repo.d.Model(blogPost).Association("Comments").Append(&comment)
	repo.d.Model(blogPost).Update("comments_count", blogPost.CommentsCount+1)
	return nil
}

// GetCommentByUserID implements domain.BlogRepository.
func (repo *blogRepo) GetCommentByUserIDRepo(blogPost *models.BlogPost, commentID uint) (models.Comment, error) {
	var comment models.Comment
	err := repo.d.Where("id = ? AND blog_post_id = ?", commentID, blogPost.ID).First(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

// GetComments implements domain.BlogRepository.
func (repo *blogRepo) GetCommentsRepo(blogPost *models.BlogPost) ([]models.Comment, error) {
	var comments []models.Comment
	err := repo.d.Where("blog_post_id = ?", blogPost.ID).Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}

// DeleteComment implements domain.BlogRepository.
func (repo *blogRepo) DeleteCommentRepo(blogPost *models.BlogPost, commentID uint) error {
	var comment models.Comment
	err := repo.d.Where("id = ? AND blog_post_id = ?", commentID, blogPost.ID).First(&comment).Error
	if err != nil {
		return err
	}
	err = repo.d.Delete(&comment).Error
	if err != nil {
		return err
	}
	repo.d.Model(blogPost).Association("Comments").Delete(&comment)
	repo.d.Model(blogPost).Update("comments_count", blogPost.CommentsCount-1)
	return nil
}

// UpdateComment implements domain.BlogRepository.
func (repo *blogRepo) UpdateCommentRepo(blogPost *models.BlogPost, comment *models.Comment) error {
	var existingComment models.Comment
	err := repo.d.Where("id = ? AND blog_post_id = ?", comment.ID, blogPost.ID).First(&existingComment).Error
	if err != nil {
		return err
	}
	err = repo.d.Model(&existingComment).Updates(comment).Error
	if err != nil {
		return err
	}
	return nil
}
