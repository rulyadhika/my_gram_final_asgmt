package handler

import (
	"log"
	"net/http"
	"strconv"

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
	result, err := s.SocialMediaService.FindAll(ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "successfully get all social media",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *SocialMediaHandlerImpl) Create(ctx *gin.Context) {
	socialMediaDto := &dto.NewSocialMediaRequest{}

	err := ctx.ShouldBindJSON(socialMediaDto)

	if err != nil {
		log.Printf("[CreateSocialMedia - Handler] err: %s \n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	// TODO change to dynamic user id
	socialMediaDto.UserId = 1

	result, err := s.SocialMediaService.Create(ctx, socialMediaDto)

	if err != nil {
		ctx.Error(err)
		return
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
	socialMediaId := ctx.Param("socialMediaId")

	socialMediaDto := &dto.UpdateSocialMediaRequest{}

	err := ctx.ShouldBindJSON(socialMediaDto)

	if err != nil {
		log.Printf("[UpdateSocialMedia - Handler] err: %s \n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	socmedId, err := strconv.Atoi(socialMediaId)
	if err != nil {
		log.Printf("[UpdateSocialMedia - Handler] err: %s \n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("param socialMediaId should be a number"))
		return
	}

	socialMediaDto.Id = uint(socmedId)

	result, err := s.SocialMediaService.Update(ctx, socialMediaDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "successfully update social media",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *SocialMediaHandlerImpl) Delete(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")

	socmedId, err := strconv.Atoi(socialMediaId)
	if err != nil {
		log.Printf("[DeleteSocialMedia - Handler] err: %s \n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("param socialMediaId should be a number"))
		return
	}

	err = s.SocialMediaService.Delete(ctx, socmedId)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "Your social media has been successfully deleted",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, response)
}
