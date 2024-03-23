package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
)

type SocialMediaRepository interface {
	FindAll(ctx *gin.Context, db *sql.DB) ([]SocialMediaUser, error)
	FindById(ctx *gin.Context, db *sql.DB, socialMediaId int) (entity.SocialMedia, error)
	Create(ctx *gin.Context, db *sql.DB, socialMedia entity.SocialMedia) (entity.SocialMedia, error)
	Update(ctx *gin.Context, db *sql.DB, socialMedia entity.SocialMedia) (entity.SocialMedia, error)
	Delete(ctx *gin.Context, db *sql.DB, socialMediaId int) error
}
