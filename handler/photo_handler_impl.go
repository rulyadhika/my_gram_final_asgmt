package handler

import (
	"github.com/gin-gonic/gin"
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
	panic("not implemented") // TODO: Implement
}

func (p *PhotoHandlerImpl) Update(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (p *PhotoHandlerImpl) Delete(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}
