package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
)

type CommentRepositoryImpl struct{}

func (c *CommentRepositoryImpl) FindAll(ctx *gin.Context, db *sql.DB) (*[]entity.Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (c *CommentRepositoryImpl) Create(ctx *gin.Context, db *sql.DB, comment *entity.Comment) (*entity.Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (c *CommentRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, comment *entity.Comment) (*entity.Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (c *CommentRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, commentId int) error {
	panic("not implemented") // TODO: Implement
}
