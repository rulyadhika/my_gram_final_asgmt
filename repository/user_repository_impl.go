package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
)

type UserRepositoryImpl struct{}

func (u *UserRepositoryImpl) Register(ctx *gin.Context, db *sql.DB, user *entity.User) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserRepositoryImpl) Login(ctx *gin.Context, db *sql.DB, user *entity.User) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, user *entity.User) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, userId int) error {
	panic("not implemented") // TODO: Implement
}
