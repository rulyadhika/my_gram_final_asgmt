package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
