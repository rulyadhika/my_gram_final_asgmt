package handler

import (
	"log"
	"net/http"

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
	panic("not implemented") // TODO: Implement
}

func (p *PhotoHandlerImpl) Create(ctx *gin.Context) {
	photoDto := &dto.NewPhotoRequest{}

	err := ctx.ShouldBindJSON(photoDto)

	if err != nil {
		log.Printf("[CreatePhoto - Handler] err: %s", err.Error())
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
	panic("not implemented") // TODO: Implement
}

func (p *PhotoHandlerImpl) Delete(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}
