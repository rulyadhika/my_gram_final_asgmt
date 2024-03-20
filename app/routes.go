package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/handler"
)

func SocialMediaRoutes(r *gin.Engine, handler handler.SocialMediaHandler) {
	socialMediaRoute := r.Group("/socialmedias")
	{
		socialMediaRoute.GET("/", handler.FindAll)
		socialMediaRoute.POST("/", handler.Create)
		socialMediaRoute.PUT("/:socialMediaId", handler.Update)
		socialMediaRoute.DELETE("/:socialMediaId", handler.Delete)
	}
}
