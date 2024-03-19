package handler

import "github.com/gin-gonic/gin"

type SocialMediaHandler interface {
	FindAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
