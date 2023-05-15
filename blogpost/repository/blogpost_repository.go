package repository

import (
	"context"
	"errors"

	"github.com/jchawla2804/go-webdev-blog/domain"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type blogPostRepo struct {
	Db *gorm.DB
}

func NewBlogPostRepo(db *gorm.DB) domain.BlogPostRepository {
	return &blogPostRepo{
		Db: db,
	}
}

func (blogPostRepo *blogPostRepo) fetchFromDb(ctx context.Context, field string, value interface{}) (domain.BlogPost, error) {

	var blogPost domain.BlogPost
	log.Info("Query DB")

	if field == "Title" {
		value = value.(string)

	} else {
		value = value.(int64)
	}

	log.Printf("%T", value)

	err := blogPostRepo.Db.WithContext(ctx).Where(field+" = ?", value).First(&blogPost).Error
	//log.Info(blogPost)
	if err != nil {
		log.Errorf("Error Occured:  %s", err.Error())
		return domain.BlogPost{}, err
	}
	if (blogPost.Title == "") || (blogPost.Id == 0) {
		log.Info("No Record found")
		return domain.BlogPost{}, errors.New("No Record found")
	}

	return blogPost, nil
}

func (blogPostRepo *blogPostRepo) CreateBlogPost(ctx context.Context, blogPost *domain.BlogPost) error {
	log.Info("Create a new blog post")
	createStatus := blogPostRepo.Db.WithContext(ctx).Create(blogPost)
	if createStatus.Error != nil {
		log.Error("Error Occured " + createStatus.Error.Error())
		return createStatus.Error
	}
	log.Info("Create Blog POST Successful")
	log.Info(blogPost.Id)
	return nil
}

func (blogPostRepo *blogPostRepo) GetBlogPostByTitle(ctx context.Context, title string) (domain.BlogPost, error) {
	log.Info("Query blogpost by title")
	blogPostResp, err := blogPostRepo.fetchFromDb(ctx, "Title", title)
	return blogPostResp, err

}

func (blogPostRepo *blogPostRepo) GetBlogPostById(ctx context.Context, id int64) (domain.BlogPost, error) {
	log.Info("Query blogpost by title")
	blogPostResp, err := blogPostRepo.fetchFromDb(ctx, "Id", id)
	return blogPostResp, err
}

func (blogPostRepo *blogPostRepo) UpdateBlogPost(ctx context.Context, blogPost *domain.BlogPost) error {
	log.Info("Update blog post")
	title := blogPost.Title
	blogPostResp, err := blogPostRepo.fetchFromDb(ctx, "Title", title)
	if err != nil {
		return err
	}

	blogPostResp.Title = blogPost.Title
	blogPostResp.Content = blogPost.Content
	blogPostResp.Description = blogPost.Description
	blogPostResp.FeatureImage = blogPost.FeatureImage
	blogPostResp.Tags = blogPost.Tags

	return blogPostRepo.Db.WithContext(ctx).Save(blogPostResp).Error

}

func (blogPostRepo *blogPostRepo) DeleteBlogPost(ctx context.Context, title string) error {
	log.Info("Update blog post")
	blogPostResp, err := blogPostRepo.fetchFromDb(ctx, "Title", title)
	if err != nil {
		return err
	}
	return blogPostRepo.Db.WithContext(ctx).Delete(blogPostResp).Error

}
