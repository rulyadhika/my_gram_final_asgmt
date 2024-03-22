package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/handler"
)

func PhotoRoutes(r *gin.Engine, handler handler.PhotoHandler) {
	photoRoute := r.Group("/photos")
	{
		photoRoute.GET("/", handler.FindAll)
		photoRoute.POST("/", handler.Create)
		photoRoute.PUT("/:photoId", handler.Update)
		photoRoute.DELETE("/:photoId", handler.Delete)
	}
}

func SocialMediaRoutes(r *gin.Engine, handler handler.SocialMediaHandler) {
	socialMediaRoute := r.Group("/socialmedias")
	{
		socialMediaRoute.GET("/", handler.FindAll)
		socialMediaRoute.POST("/", handler.Create)
		socialMediaRoute.PUT("/:socialMediaId", handler.Update)
		socialMediaRoute.DELETE("/:socialMediaId", handler.Delete)
	}
}

func CommentRoutes(r *gin.Engine, handler handler.CommentHandler) {
	commentRoute := r.Group("/comments")
	{
		commentRoute.GET("/", handler.FindAll)
		commentRoute.POST("/", handler.Create)
		commentRoute.PUT("/:commentId", handler.Update)
		commentRoute.DELETE("/:commentId", handler.Delete)
	}
}

func UserRoutes(r *gin.Engine, handler handler.UserHandler) {
	userRoute := r.Group("/users")
	{
		userRoute.POST("/register", handler.Register)
		userRoute.POST("/login", handler.Login)
		userRoute.PUT("/", handler.Update)
		userRoute.DELETE("/", handler.Delete)
	}
}
