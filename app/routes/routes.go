package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/handler"
	"github.com/rulyadhika/my_gram_final_asgmt/middleware"
)

func PhotoRoutes(r *gin.Engine, handler handler.PhotoHandler, authMiddleware middleware.AuthMiddleware) {
	photoRoute := r.Group("/photos")
	{
		photoRoute.Use(authMiddleware.Authentication())

		photoRoute.GET("/", handler.FindAll)
		photoRoute.POST("/", handler.Create)
		photoRoute.PUT("/:photoId", handler.Update)
		photoRoute.DELETE("/:photoId", handler.Delete)
	}
}

func SocialMediaRoutes(r *gin.Engine, handler handler.SocialMediaHandler, authMiddleware middleware.AuthMiddleware) {
	socialMediaRoute := r.Group("/socialmedias")
	{
		socialMediaRoute.Use(authMiddleware.Authentication())

		socialMediaRoute.GET("/", handler.FindAll)
		socialMediaRoute.POST("/", handler.Create)
		socialMediaRoute.PUT("/:socialMediaId", handler.Update)
		socialMediaRoute.DELETE("/:socialMediaId", handler.Delete)
	}
}

func CommentRoutes(r *gin.Engine, handler handler.CommentHandler, authMiddleware middleware.AuthMiddleware) {
	commentRoute := r.Group("/comments")
	{
		commentRoute.Use(authMiddleware.Authentication())

		commentRoute.GET("/", handler.FindAll)
		commentRoute.POST("/", handler.Create)
		commentRoute.PUT("/:commentId", handler.Update)
		commentRoute.DELETE("/:commentId", handler.Delete)
	}
}

func UserRoutes(r *gin.Engine, handler handler.UserHandler, authMiddleware middleware.AuthMiddleware) {
	userRoute := r.Group("/users")
	{
		userRoute.POST("/register", handler.Register)
		userRoute.POST("/login", handler.Login)

		userRoute.Use(authMiddleware.Authentication())
		userRoute.PUT("/", handler.Update)
		userRoute.DELETE("/", handler.Delete)
	}
}
