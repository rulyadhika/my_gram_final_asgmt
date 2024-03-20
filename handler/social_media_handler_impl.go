package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
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
	socialMediaDto := &dto.NewSocialMediaRequest{}

	err := ctx.ShouldBindJSON(socialMediaDto)

	if err != nil {
		ctx.Error(errs.NewUnprocessableEntityError("invalid json request body"))
	}

	// TODO change to dynamic user id
	socialMediaDto.UserId = 1

	result, err := s.SocialMediaService.Create(ctx, socialMediaDto)

	if err != nil {
		ctx.Error(err)
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusCreated),
		Code:    http.StatusCreated,
		Message: "successfully created new social media",
		Data:    result,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (s *SocialMediaHandlerImpl) Update(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaHandlerImpl) Delete(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}
