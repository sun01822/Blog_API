package models

import (
	"gorm.io/gorm"
	"time"
)

type BlogPost struct {
	gorm.Model
	UserID        uint       `json:"user_id"`
	Title         string     `json:"title"`
	ContentText   string     `json:"content_text"`
	PhotoURL      string     `json:"photo_url"`
	Description   string     `json:"description"`
	Category      string     `json:"category"`
	Comments      []Comment  `json:"comments"`
	CommentsCount int        `json:"comments_count"`
	Likes         []Like     `json:"likes"`
	LikesCount    int        `json:"likes_count"`
	Views         int        `json:"views"`
	IsPublished   bool       `json:"is_published"`
	PublishedAt   *time.Time `json:"published_at"`
}

type Comment struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	BlogPost  BlogPost
	BlogPostID uint   `json:"blog_post_id"`
	Content   string `json:"content"`
}

type Like struct { 
	gorm.Model
	UserID    uint   `json:"user_id"`
	BlogPost  BlogPost
	BlogPostID uint   `json:"blog_post_id"`
}

