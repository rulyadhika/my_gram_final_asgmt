package repository

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
)

type PhotoRepositoryImpl struct{}

func NewPhotoRepositoryImpl() PhotoRepository {
	return &PhotoRepositoryImpl{}
}

func (p *PhotoRepositoryImpl) FindAll(ctx *gin.Context, db *sql.DB) (*[]entity.Photo, error) {
	panic("not implemented") // TODO: Implement
}

func (p *PhotoRepositoryImpl) Create(ctx *gin.Context, db *sql.DB, photo entity.Photo) (entity.Photo, error) {
	sqlQuery := `INSERT INTO photos(title, caption, photo_url, user_id) VALUES($1, $2, $3, $4) RETURNING id, created_at`

	err := db.QueryRowContext(ctx, sqlQuery, photo.Title, photo.Caption, photo.PhotoUrl, photo.UserId).Scan(&photo.Id, &photo.CreatedAt)

	if err != nil {
		log.Printf("[CreatePhoto - Repo] err: %s", err.Error())
		return photo, errs.NewInternalServerError("something went wrong")
	}

	return photo, nil
}

func (p *PhotoRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, photo entity.Photo) (entity.Photo, error) {
	panic("not implemented") // TODO: Implement
}

func (p *PhotoRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, photoId int) error {
	panic("not implemented") // TODO: Implement
}
