package domain

import (
	"context"

	"github.com/lib/pq"
)

type BlogPost struct {
	Id           int64          `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Title        string         `json:"title" gorm:"unique" `
	FeatureImage string         `json:"featureImage"`
	Description  string         `json:"description"`
	Content      string         `json:"content"`
	Tags         pq.StringArray `gorm:"type:text[]" json:"tags"`
}

type BlogPostRepository interface {
	CreateBlogPost(ctx context.Context, blog *BlogPost) error
	GetBlogPostByTitle(ctx context.Context, title string) (BlogPost, error)
	GetBlogPostById(ctx context.Context, id int64) (BlogPost, error)
	UpdateBlogPost(ctx context.Context, blogPost *BlogPost) error
	DeleteBlogPost(ctx context.Context, title string) error
}

type BlogPostUsecase interface {
	CreateBlogPost(ctx context.Context, blog *BlogPost) error
	GetBlogPostByTitle(ctx context.Context, title string) (BlogPost, error)
	UpdateBlogPost(ctx context.Context, blogPost *BlogPost) error
	DeleteBlogPost(ctx context.Context, title string) error
	GetBlogPostById(ctx context.Context, id int64) (BlogPost, error)
}
