package types

import (
	"github.com/go-ozzo/ozzo-validation"
)

type BlogPostRequest struct {
	Title       string `json:"title"`
	ContentText string `json:"content_text"`
	PhotoURL    string `json:"photo_url"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IsPublished bool   `json:"is_published"`
}

func (blogPost BlogPostRequest) Validate() error {
	return validation.ValidateStruct(&blogPost,
		validation.Field(&blogPost.Title, validation.Required, validation.Length(10, 255)),
		validation.Field(&blogPost.Category, validation.Required, validation.Length(3, 100)),
	)
}

type UpdateBlogPostRequest struct {
	Title       string `json:"title"`
	ContentText string `json:"content_text"`
	PhotoURL    string `json:"photo_url"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IsPublished bool   `json:"is_published"`
}

type Comment struct {
	Content string `json:"content"`
}

func (comment Comment) Validate() error {
	return validation.ValidateStruct(&comment,
		validation.Field(&comment.Content, validation.Required, validation.Length(2, 500)),
	)
}

func (blogPost UpdateBlogPostRequest) Validate() error {
	return validation.ValidateStruct(&blogPost,
		validation.Field(&blogPost.Title, validation.Required, validation.Length(10, 255)),
		validation.Field(&blogPost.Category, validation.Required, validation.Length(3, 100)),
	)
}

type BlogResp struct {
	ID             string         `json:"id,omitempty"`
	UserID         string         `json:"user_id,omitempty"`
	Title          string         `json:"title,omitempty"`
	ContentText    string         `json:"content_text,omitempty"`
	PhotoURL       string         `json:"photo_url,omitempty"`
	Description    string         `json:"description,omitempty"`
	Category       string         `json:"category"`
	Comments       []CommentResp  `json:"comments"`
	CommentsCount  uint           `json:"comments_count"`
	ReactionsCount uint           `json:"reactions_count"`
	Reactions      []ReactionResp `json:"reactions"`
	Views          uint           `json:"views"`
	IsPublished    bool           `json:"is_published"`
	PublishedAt    string         `json:"published_at"`
	CreatedAt      string         `json:"created_at,omitempty"`
	UpdatedAt      string         `json:"updated_at,omitempty"`
	DeletedAt      string         `json:"deleted_at,omitempty"`
}

type ReactionResp struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	BlogPostID string `json:"blog_post_id"`
	Type       uint64 `json:"type"`
}

type CommentResp struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	BlogPostID string `json:"blog_post_id"`
	Content    string `json:"content"`
}
