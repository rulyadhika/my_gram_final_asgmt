package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
)

type UserRepository interface {
	Register(ctx *gin.Context, db *sql.DB, user *entity.User) (*entity.User, error)
	Login(ctx *gin.Context, db *sql.DB, user *entity.User) (*entity.User, error)
	Update(ctx *gin.Context, db *sql.DB, user *entity.User) (*entity.User, error)
	Delete(ctx *gin.Context, db *sql.DB, userId int) error
}
