package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
)

type CommentService interface {
	FindAll(ctx *gin.Context) (*[]dto.CommentResponse, error)
	Create(ctx *gin.Context, commentDto *dto.NewCommentRequest) (*dto.NewCommentResponse, error)
	Update(ctx *gin.Context, commentDto *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, error)
	Delete(ctx *gin.Context, commentId int) error
}
