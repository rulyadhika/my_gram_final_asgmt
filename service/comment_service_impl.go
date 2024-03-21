package service

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/helper"
	"github.com/rulyadhika/my_gram_final_asgmt/repository"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCommentServiceImpl(commentRepository repository.CommentRepository, db *sql.DB, validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (c *CommentServiceImpl) FindAll(ctx *gin.Context) (*[]dto.CommentResponse, error) {
	result, err := c.CommentRepository.FindAll(ctx, c.DB)

	if err != nil {
		return &[]dto.CommentResponse{}, err
	}

	return helper.ToCommentsResponse(result), nil
}

func (c *CommentServiceImpl) Create(ctx *gin.Context, commentDto *dto.NewCommentRequest) (*dto.NewCommentResponse, error) {
	validationErr := c.Validate.Struct(commentDto)

	if validationErr != nil {
		return &dto.NewCommentResponse{}, validationErr
	}

	photoId, err := (commentDto.PhotoId).Float64()

	if err != nil {
		return &dto.NewCommentResponse{}, err
	}

	comment := entity.Comment{
		PhotoId: uint(photoId),
		UserId:  commentDto.UserId,
		Message: commentDto.Message,
	}

	result, err := c.CommentRepository.Create(ctx, c.DB, comment)

	if err != nil {
		return &dto.NewCommentResponse{}, err
	}

	return helper.ToNewCommentResponse(result), nil
}

func (c *CommentServiceImpl) Update(ctx *gin.Context, commentDto *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, error) {
	validationErr := c.Validate.Struct(commentDto)

	if validationErr != nil {
		return &dto.UpdateCommentResponse{}, validationErr
	}

	comment := entity.Comment{
		Id:      commentDto.Id,
		Message: commentDto.Message,
	}

	result, err := c.CommentRepository.Update(ctx, c.DB, comment)

	if err != nil {
		return &dto.UpdateCommentResponse{}, err
	}

	return helper.ToUpdateCommentResponse(result), nil
}

func (c *CommentServiceImpl) Delete(ctx *gin.Context, commentId int) error {
	err := c.CommentRepository.Delete(ctx, c.DB, commentId)

	if err != nil {
		return err
	}

	return nil
}
