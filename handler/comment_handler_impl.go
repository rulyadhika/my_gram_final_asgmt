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

type CommentHandlerImpl struct {
	CommentService service.CommentService
}

func NewCommentHandlerImpl(commentService service.CommentService) CommentHandler {
	return &CommentHandlerImpl{CommentService: commentService}
}

func (c *CommentHandlerImpl) FindAll(ctx *gin.Context) {
	result, err := c.CommentService.FindAll(ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "successfully get all comments",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *CommentHandlerImpl) Create(ctx *gin.Context) {
	commentDto := &dto.NewCommentRequest{}

	err := ctx.ShouldBindJSON(commentDto)

	if err != nil {
		log.Printf("[CreateComment - handler] err: %s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	// TODO userId should be dynamic from context
	userId := 2
	commentDto.UserId = uint(userId)

	result, err := c.CommentService.Create(ctx, commentDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusCreated),
		Code:    http.StatusCreated,
		Message: "successfully created comment",
		Data:    result,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *CommentHandlerImpl) Update(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	id, err := strconv.Atoi(commentId)

	if err != nil {
		log.Printf("[UpdateComment - Handler] err:%s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("param commentId must be a number"))
		return
	}

	commentDto := &dto.UpdateCommentRequest{}
	err = ctx.ShouldBindJSON(commentDto)

	if err != nil {
		log.Printf("[UpdateComment - Handler] err:%s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	commentDto.Id = uint(id)

	result, err := c.CommentService.Update(ctx, commentDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "successfully update comment",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *CommentHandlerImpl) Delete(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	id, err := strconv.Atoi(commentId)

	if err != nil {
		log.Printf("[UpdateComment - Handler] err:%s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("param commentId must be a number"))
		return
	}

	err = c.CommentService.Delete(ctx, id)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "successfully delete comment",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, response)
}
