package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
)

type PhotoService interface {
	FindAll(ctx *gin.Context) (*[]dto.PhotoResponse, error)
	Create(ctx *gin.Context, photoDto *dto.NewPhotoRequest) (*dto.NewPhotoResponse, error)
	Update(ctx *gin.Context, photoDto *dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, error)
	Delete(ctx *gin.Context, photoId int) error
}
