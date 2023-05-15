package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jchawla2804/go-webdev-blog/domain"
	"github.com/labstack/gommon/log"
)

type blogPostHandler struct {
	blogPostUsecase domain.BlogPostUsecase
}

func NewBlogPostHandler(blogPostUseCase domain.BlogPostUsecase, ghEngine *gin.Engine) {
	blogPostHandler := &blogPostHandler{
		blogPostUsecase: blogPostUseCase,
	}
	ghEngine.GET("/api/blog", blogPostHandler.GetBlogPostByTitle)
	ghEngine.POST("/api/blog", blogPostHandler.CreateBlogPost)
	ghEngine.DELETE("/api/blog", blogPostHandler.DeleteBlogPostByTitle)
	ghEngine.PUT("/api/blog", blogPostHandler.UpdateBlogPost)
	ghEngine.GET("/api/blog/:id", blogPostHandler.GetBlogPostById)
}

func (bph *blogPostHandler) GetBlogPostById(gh *gin.Context) {
	id := gh.Param("id")
	blogPostId, _ := strconv.Atoi(id)
	blogPostResp, err := bph.blogPostUsecase.GetBlogPostById(gh.Request.Context(), int64(blogPostId))
	if err != nil {
		if err.Error() == "No Record found" {
			gh.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		gh.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	gh.JSON(http.StatusOK, blogPostResp)
	return
}

func (bph *blogPostHandler) GetBlogPostByTitle(gh *gin.Context) {
	log.Info("Get Blog post by title request")
	title := gh.Query("title")
	log.Info(title)
	if title == "" {
		gh.JSON(http.StatusBadRequest, gin.H{
			"error": "Please pass blog post title",
		})
		return
	}

	parentContext := gh.Request.Context()
	blogPostResp, err := bph.blogPostUsecase.GetBlogPostByTitle(parentContext, title)

	if err != nil {
		if err.Error() == "No Record found" {
			gh.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		gh.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		gh.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	gh.JSON(http.StatusOK, blogPostResp)
	return
}

func (bph *blogPostHandler) CreateBlogPost(gh *gin.Context) {
	var blogPost domain.BlogPost
	err := gh.Bind(&blogPost)
	if err != nil {
		gh.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = bph.blogPostUsecase.CreateBlogPost(gh.Request.Context(), &blogPost)
	if err != nil {
		gh.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	gh.JSON(http.StatusCreated, gin.H{
		"status": "Record has been created",
	})
	return
}

func (bph *blogPostHandler) DeleteBlogPostByTitle(gh *gin.Context) {
	title := gh.Query("title")
	if title == "" {
		gh.JSON(http.StatusBadRequest, gin.H{
			"error": "Title is not provided",
		})
		return
	}

	err := bph.blogPostUsecase.DeleteBlogPost(gh.Request.Context(), title)
	if err != nil {
		gh.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	gh.JSON(http.StatusOK, gin.H{
		"status": "Record with " + title + " has been deleted",
	})
	return
}

func (bph *blogPostHandler) UpdateBlogPost(gh *gin.Context) {
	var blogPost domain.BlogPost
	err := gh.Bind(&blogPost)
	if err != nil {
		gh.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = bph.blogPostUsecase.UpdateBlogPost(gh.Request.Context(), &blogPost)
	if err != nil {
		gh.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	gh.JSON(http.StatusNoContent, gin.H{
		"status": "Record has been updated",
	})
	return
}
