package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
)

type PhotoRepositoryImpl struct{}

func (p *PhotoRepositoryImpl) FindAll(ctx *gin.Context, db *sql.DB) (*[]entity.Photo, error) {
	panic("not implemented") // TODO: Implement
}

func (p *PhotoRepositoryImpl) Create(ctx *gin.Context, db *sql.DB, photo *entity.Photo) (*entity.Photo, error) {
	panic("not implemented") // TODO: Implement
}

func (p *PhotoRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, photo *entity.Photo) (*entity.Photo, error) {
	panic("not implemented") // TODO: Implement
}

func (p *PhotoRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, photoId int) error {
	panic("not implemented") // TODO: Implement
}
