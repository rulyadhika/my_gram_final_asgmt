package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
)

type PhotoRepository interface {
	FindAll(ctx *gin.Context, db *sql.DB) ([]PhotoUser, error)
	FindById(ctx *gin.Context, db *sql.DB, photoId int) (entity.Photo, error)
	Create(ctx *gin.Context, db *sql.DB, photo entity.Photo) (entity.Photo, error)
	Update(ctx *gin.Context, db *sql.DB, photo entity.Photo) (entity.Photo, error)
	Delete(ctx *gin.Context, db *sql.DB, photoId int) error
}
