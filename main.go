package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jchawla2804/go-webdev-blog/blogpost/delivery/handler"
	"github.com/jchawla2804/go-webdev-blog/blogpost/repository"
	"github.com/jchawla2804/go-webdev-blog/blogpost/usecase"
	"github.com/jchawla2804/go-webdev-blog/domain"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	ghEngine := gin.Default()
	dbPort, _ := strconv.Atoi(os.Getenv("db.port"))
	dbConfig := domain.DBConfig{
		DBHost:     os.Getenv("db.host"),
		DBPort:     dbPort,
		DBUsername: os.Getenv("db.user"),
		DBPassword: os.Getenv("db.pass"),
		DBName:     os.Getenv("db.name"),
	}

	gormDb, err := domain.SetUpDBConnection(dbConfig)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		db, err := gormDb.DB()
		if err != nil {
			log.Fatal(err)
		}

		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}

	}()
	gormDb.AutoMigrate(&domain.BlogPost{})

	_blogPostRepo := repository.NewBlogPostRepo(gormDb)
	_blogPostUseCase := usecase.NewBlogPostUsecase(_blogPostRepo, 10*time.Second)
	handler.NewBlogPostHandler(_blogPostUseCase, ghEngine)
	log.Fatal(ghEngine.Run(":8071"))
}
