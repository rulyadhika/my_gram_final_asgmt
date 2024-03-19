package handler

import "github.com/gin-gonic/gin"

type CommentHandler interface {
	FindAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
