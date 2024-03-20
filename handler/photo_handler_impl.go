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

type PhotoHandlerImpl struct {
	PhotoService service.PhotoService
}

func NewPhotoHandlerImpl(photoService service.PhotoService) PhotoHandler {
	return &PhotoHandlerImpl{PhotoService: photoService}
}

func (p *PhotoHandlerImpl) FindAll(ctx *gin.Context) {
	result, err := p.PhotoService.FindAll(ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "successfully get all photos",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (p *PhotoHandlerImpl) Create(ctx *gin.Context) {
	photoDto := &dto.NewPhotoRequest{}

	err := ctx.ShouldBindJSON(photoDto)

	if err != nil {
		log.Printf("[CreatePhoto - Handler] err: %s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	// TODO	user id should be dynamic from context
	photoDto.UserId = 1

	result, err := p.PhotoService.Create(ctx, photoDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusCreated),
		Code:    http.StatusOK,
		Message: "successfully create new photo",
		Data:    result,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (p *PhotoHandlerImpl) Update(ctx *gin.Context) {
	photoId := ctx.Param("photoId")
	id, err := strconv.Atoi(photoId)

	if err != nil {
		log.Printf("[UpdatePhoto - handler] err:%s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("photoId params must be a number"))
		return
	}

	photoDto := &dto.UpdatePhotoRequest{}
	err = ctx.ShouldBindJSON(photoDto)

	if err != nil {
		log.Printf("[UpdatePhoto - handler] err:%s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	photoDto.Id = uint(id)

	result, err := p.PhotoService.Update(ctx, photoDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "successfuly update photo",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (p *PhotoHandlerImpl) Delete(ctx *gin.Context) {
	photoId := ctx.Param("photoId")
	id, err := strconv.Atoi(photoId)

	if err != nil {
		log.Printf("[DeletePhoto - handler] err:%s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("photoId param must be a number"))
		return
	}

	err = p.PhotoService.Delete(ctx, id)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "successfully delete photo",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, response)
}
