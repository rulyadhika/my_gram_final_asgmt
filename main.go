package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rulyadhika/my_gram_final_asgmt/app"
	"github.com/rulyadhika/my_gram_final_asgmt/handler"
	"github.com/rulyadhika/my_gram_final_asgmt/middleware"
	"github.com/rulyadhika/my_gram_final_asgmt/repository"
	"github.com/rulyadhika/my_gram_final_asgmt/service"
)

func main() {
	appConfig := app.GetAppConfig()
	db := app.InitDB()
	validate := validator.New()
	router := gin.Default()
	router.Use(middleware.ErrorHandlerMiddleware())

	socialMediaRepository := repository.NewSocialMediaRepositoryImpl()
	socialMediaService := service.NewSocialMediaServiceImpl(socialMediaRepository, db, validate)
	socialMediaHandler := handler.NewSocialMediaHandlerImpl(socialMediaService)

	app.SocialMediaRoutes(router, socialMediaHandler)

	router.Run(":" + appConfig.SERVER_PORT)
}
