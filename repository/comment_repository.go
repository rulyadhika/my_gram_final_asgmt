package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
)

type CommentRepository interface {
	FindAll(ctx *gin.Context, db *sql.DB) ([]CommentPhotoUser, error)
	FindById(ctx *gin.Context, db *sql.DB, commentId int) (entity.Comment, error)
	Create(ctx *gin.Context, db *sql.DB, comment entity.Comment) (entity.Comment, error)
	Update(ctx *gin.Context, db *sql.DB, comment entity.Comment) (entity.Comment, error)
	Delete(ctx *gin.Context, db *sql.DB, commentId int) error
}
