package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/service"
)

type SocialMediaHandlerImpl struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaHandlerImpl(socialMediaService service.SocialMediaService) SocialMediaHandler {
	return &SocialMediaHandlerImpl{SocialMediaService: socialMediaService}
}

func (s *SocialMediaHandlerImpl) FindAll(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaHandlerImpl) Create(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaHandlerImpl) Update(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaHandlerImpl) Delete(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}
