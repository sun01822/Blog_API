package types

import (
	"github.com/go-ozzo/ozzo-validation"
)

type BlogPostRequest struct {
	Title         string `json:"title"`
	ContentText   string `json:"content_text"`
	PhotoURL      string `json:"photo_url"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	IsPublished   bool   `json:"is_published"`
}

func (blogPost BlogPostRequest) Validate() error {
	return validation.ValidateStruct(&blogPost,
		validation.Field(&blogPost.Title, validation.Required, validation.Length(10, 100)),
		validation.Field(&blogPost.Category, validation.Required, validation.Length(3, 100)),
		validation.Field(&blogPost.IsPublished, validation.Required),
	)
}


type UpdateBlogPostRequest struct {
	Title         string `json:"title"`
	ContentText   string `json:"content_text"`
	PhotoURL      string `json:"photo_url"`
	Description   string `json:"description"`
	Category      string `json:"category"`
}


type Comment struct {
	Content   string `json:"content"`
}

func (comment Comment) Validate() error {
	return validation.ValidateStruct(&comment,
		validation.Field(&comment.Content, validation.Required, validation.Length(2, 500)),
	)
}