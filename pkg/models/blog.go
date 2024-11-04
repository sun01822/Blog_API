package models

import (
	"gorm.io/gorm"
	"time"
)

type BlogPost struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	UserID         string         `json:"user_id"`
	Title          string         `json:"title" gorm:"unique"`
	ContentText    string         `json:"content_text"`
	PhotoURL       string         `json:"photo_url"`
	Description    string         `json:"description"`
	Category       string         `json:"category"`
	Comments       []Comment      `json:"comments"`
	CommentsCount  uint           `json:"comments_count"`
	Reactions      []Reaction     `json:"reactions"`
	ReactionsCount uint           `json:"reactions_count"`
	Views          uint           `json:"views"`
	IsPublished    bool           `json:"is_published"`
	PublishedAt    time.Time      `json:"published_at"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Comment struct {
	ID         string         `json:"id" gorm:"primaryKey"`
	UserID     string         `json:"user_id" gorm:"size:255"`
	BlogPostID string         `json:"blog_post_id" gorm:"size:255"`
	Content    string         `json:"content"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Reaction struct {
	ID         string         `json:"id" gorm:"primaryKey"`
	UserID     string         `json:"user_id" gorm:"size:255"`
	BlogPostID string         `json:"blog_post_id" gorm:"size:255"`
	Type       uint           `json:"type"` // 1: like, 2: love, 3: care, 4: haha, 5: wow, 6: sad, 7: angry
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
