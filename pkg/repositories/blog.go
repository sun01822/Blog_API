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
func (repo *blogRepo) CreateBlogPost(blogPost *models.BlogPost) error {
	err := repo.d.Create(blogPost).Error
	if err != nil {
		return err
	}
	return nil
}

// GetBlogPost implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPost(id uint) (models.BlogPost, error) {
	var blogPost models.BlogPost
	err := repo.d.Where("id = ?", id).First(&blogPost).Error
	if err != nil {
		return blogPost, err
	}
	return blogPost, nil
}

// GetBlogPosts implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPosts() ([]models.BlogPost, error) {
	var blogPosts []models.BlogPost
	err := repo.d.Find(&blogPosts).Error
	if err != nil {
		return blogPosts, err
	}
	return blogPosts, nil
}

// GetBlogPosts implements domain.BlogRepository.
func (repo *blogRepo) GetBlogPostsOfUser(userID uint) ([]models.BlogPost, error) {
	var blogPosts []models.BlogPost
	err := repo.d.Where("user_id = ?", userID).Find(&blogPosts).Error
	if err != nil {
		return blogPosts, err
	}
	return blogPosts, nil
}

// UpdateBlogPost implements domain.BlogRepository.
func (repo *blogRepo) UpdateBlogPost(blogPost *models.BlogPost) error {
	err := repo.d.Save(blogPost).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteBlogPost implements domain.BlogRepository.
func (repo *blogRepo) DeleteBlogPost(id uint) error {
	err := repo.d.Where("id = ?", id).Delete(&models.BlogPost{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *blogRepo) AddAndRemoveLike(blogPost *models.BlogPost, userID uint) (string, error) {
	// check if the user has already liked the post
	// if yes, remove the like
	// if no, add the like
	// return the updated blogPost
	var like models.Like
	err := repo.d.Where("user_id = ? AND blog_post_id = ?", userID, blogPost.ID).First(&like).Error
	if err != nil {
		// user has not liked the post
		like = models.Like{
			UserID: userID,
			BlogPostID: blogPost.ID,
			BlogPost: *blogPost,
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
func (repo *blogRepo) AddComment(blogPost *models.BlogPost, comment *models.Comment) error {
	err := repo.d.Model(blogPost).Association("Comments").Append(comment)
	if err != nil {
		return err
	}
	return nil
}

// GetCommentByUserID implements domain.BlogRepository.
func (repo *blogRepo) GetCommentByUserID(blogPost *models.BlogPost, commentID uint) (models.Comment, error) {
	var comment models.Comment
	err := repo.d.Model(blogPost).Association("Comments").Find(&comment, "id = ?", commentID)
	if err != nil {
		return comment, err
	}
	return comment, nil
}


// GetComments implements domain.BlogRepository.
func (repo *blogRepo) GetComments(blogPost *models.BlogPost) ([]models.Comment, error) {
	var comments []models.Comment
	err := repo.d.Model(blogPost).Association("Comments").Find(&comments)
	if err != nil {
		return comments, err
	}
	return comments, nil
}

// DeleteComment implements domain.BlogRepository.
func (repo *blogRepo) DeleteComment(blogPost *models.BlogPost, commentID uint) error {
	err := repo.d.Model(blogPost).Association("Comments").Delete(&models.Comment{Model: gorm.Model{ID: commentID}})
	if err != nil {
		return err
	}
	return nil
}

// UpdateComment implements domain.BlogRepository.
func (repo *blogRepo) UpdateComment(blogPost *models.BlogPost, comment *models.Comment) error {
	err := repo.d.Model(blogPost).Association("Comments").Replace(comment)
	if err != nil {
		return err
	}
	return nil
}