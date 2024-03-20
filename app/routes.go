package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/handler"
)

func SocialMediaRoutes(r *gin.Engine, handler handler.SocialMediaHandler) {
	socialMediaRoute := r.Group("/social-medias")
	{
		socialMediaRoute.POST("/", handler.Create)
	}
}
