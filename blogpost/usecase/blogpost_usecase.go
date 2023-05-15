package usecase

import (
	"context"
	"time"

	"github.com/jchawla2804/go-webdev-blog/domain"
)

type blogPostUsecase struct {
	blogPostRepo domain.BlogPostRepository
	timeout      time.Duration
}

func NewBlogPostUsecase(blogRepo domain.BlogPostRepository, timeout time.Duration) domain.BlogPostUsecase {
	return &blogPostUsecase{
		blogPostRepo: blogRepo,
		timeout:      timeout,
	}
}

func (bpu *blogPostUsecase) CreateBlogPost(ctx context.Context, blog *domain.BlogPost) error {
	timeoutContext, cancel := context.WithTimeout(ctx, bpu.timeout)
	defer cancel()
	return bpu.blogPostRepo.CreateBlogPost(timeoutContext, blog)
}
func (bpu *blogPostUsecase) GetBlogPostByTitle(ctx context.Context, title string) (domain.BlogPost, error) {
	timeoutContext, cancel := context.WithTimeout(ctx, bpu.timeout)
	defer cancel()
	return bpu.blogPostRepo.GetBlogPostByTitle(timeoutContext, title)
}
func (bpu *blogPostUsecase) UpdateBlogPost(ctx context.Context, blogPost *domain.BlogPost) error {
	timeoutContext, cancel := context.WithTimeout(ctx, bpu.timeout)
	defer cancel()
	return bpu.blogPostRepo.UpdateBlogPost(timeoutContext, blogPost)

}
func (bpu *blogPostUsecase) DeleteBlogPost(ctx context.Context, title string) error {
	timeoutContext, cancel := context.WithTimeout(ctx, bpu.timeout)
	defer cancel()
	return bpu.blogPostRepo.DeleteBlogPost(timeoutContext, title)
}

func (bpu *blogPostUsecase) GetBlogPostById(ctx context.Context, id int64) (domain.BlogPost, error) {
	timeoutContext, cancel := context.WithTimeout(ctx, bpu.timeout)
	defer cancel()
	return bpu.blogPostRepo.GetBlogPostById(timeoutContext, id)
}
