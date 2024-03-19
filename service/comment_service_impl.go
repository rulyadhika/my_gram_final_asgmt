package service

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
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
	panic("not implemented") // TODO: Implement
}

func (c *CommentServiceImpl) Create(ctx *gin.Context, commentDto *dto.NewCommentRequest) (*dto.NewCommentResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (c *CommentServiceImpl) Update(ctx *gin.Context, commentDto *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (c *CommentServiceImpl) Delete(ctx *gin.Context, commentId int) error {
	panic("not implemented") // TODO: Implement
}
