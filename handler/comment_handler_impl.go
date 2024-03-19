package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/service"
)

type CommentHandlerImpl struct {
	CommentService service.CommentService
}

func NewCommentHandlerImpl(commentService service.CommentService) CommentHandler {
	return &CommentHandlerImpl{CommentService: commentService}
}

func (c *CommentHandlerImpl) FindAll(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *CommentHandlerImpl) Create(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *CommentHandlerImpl) Update(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *CommentHandlerImpl) Delete(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}
