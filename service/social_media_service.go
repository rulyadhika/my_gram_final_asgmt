package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
)

type SocialMediaService interface {
	FindAll(ctx *gin.Context) (*[]dto.SocialMediaResponse, error)
	Create(ctx *gin.Context, socialMediaDto *dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, error)
	Update(ctx *gin.Context, socialMediaDto *dto.UpdateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, error)
	Delete(ctx *gin.Context, socialMediaId int) error
}
