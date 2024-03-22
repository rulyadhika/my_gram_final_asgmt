package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rulyadhika/my_gram_final_asgmt/app/config"
	"github.com/rulyadhika/my_gram_final_asgmt/app/database"
	"github.com/rulyadhika/my_gram_final_asgmt/app/routes"
	"github.com/rulyadhika/my_gram_final_asgmt/handler"
	"github.com/rulyadhika/my_gram_final_asgmt/middleware"
	"github.com/rulyadhika/my_gram_final_asgmt/repository"
	"github.com/rulyadhika/my_gram_final_asgmt/service"
)

func main() {
	appConfig := config.GetAppConfig()
	db := database.InitDB()
	validate := validator.New()
	router := gin.Default()
	router.Use(middleware.ErrorHandlerMiddleware())

	// social media
	socialMediaRepository := repository.NewSocialMediaRepositoryImpl()
	socialMediaService := service.NewSocialMediaServiceImpl(socialMediaRepository, db, validate)
	socialMediaHandler := handler.NewSocialMediaHandlerImpl(socialMediaService)
	// social media

	// photo
	photoRepository := repository.NewPhotoRepositoryImpl()
	photoService := service.NewPhotoServiceImpl(photoRepository, db, validate)
	photoHandler := handler.NewPhotoHandlerImpl(photoService)
	// photo

	// comment
	commentRepository := repository.NewCommentRepositoryImpl()
	commentService := service.NewCommentServiceImpl(commentRepository, db, validate)
	commentHandler := handler.NewCommentHandlerImpl(commentService)
	// comment

	// routes
	routes.SocialMediaRoutes(router, socialMediaHandler)
	routes.PhotoRoutes(router, photoHandler)
	routes.CommentRoutes(router, commentHandler)
	// end of routes

	router.Run(":" + appConfig.SERVER_PORT)
}
