package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
)

type PhotoRepositoryImpl struct{}

func NewPhotoRepositoryImpl() PhotoRepository {
	return &PhotoRepositoryImpl{}
}

func (p *PhotoRepositoryImpl) FindAll(ctx *gin.Context, db *sql.DB) ([]PhotoUser, error) {
	sqlQuery := `SELECT photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at,
	photos.updated_at, users.id, users.email, users.username FROM photos JOIN users ON photos.user_id=users.id`

	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		log.Printf("[FindAllPhoto - Repo] err: %s\n", err.Error())
	}
	defer rows.Close()

	photosUser := []PhotoUser{}

	for rows.Next() {
		photoUser := PhotoUser{}

		err := rows.Scan(&photoUser.Photo.Id, &photoUser.Title, &photoUser.Caption, &photoUser.PhotoUrl, &photoUser.UserId, &photoUser.Photo.CreatedAt, &photoUser.Photo.UpdatedAt, &photoUser.User.Id, &photoUser.Email, &photoUser.Username)
		if err != nil {
			log.Printf("[FindAllPhoto - Repo] err: %s\n", err.Error())
			return photosUser, errs.NewInternalServerError("something went wrong")
		}

		photosUser = append(photosUser, photoUser)
	}

	return photosUser, nil
}

func (p *PhotoRepositoryImpl) FindById(ctx *gin.Context, db *sql.DB, photoId int) (entity.Photo, error) {
	sqlQuery := `SELECT id, title, caption, photo_url, user_id, created_at, updated_at FROM photos WHERE id=$1`

	photo := entity.Photo{}

	err := db.QueryRowContext(ctx, sqlQuery, photoId).Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt)

	if err != nil {
		log.Printf("[FindPhotoById - Repo] err: %s \n", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return photo, errs.NewNotFoundError(fmt.Sprintf("social media with id:%v not found", photoId))
		}

		return photo, errs.NewInternalServerError("something went wrong")
	}

	return photo, nil
}

func (p *PhotoRepositoryImpl) Create(ctx *gin.Context, db *sql.DB, photo entity.Photo) (entity.Photo, error) {
	sqlQuery := `INSERT INTO photos(title, caption, photo_url, user_id) VALUES($1, $2, $3, $4) RETURNING id, created_at`

	err := db.QueryRowContext(ctx, sqlQuery, photo.Title, photo.Caption, photo.PhotoUrl, photo.UserId).Scan(&photo.Id, &photo.CreatedAt)

	if err != nil {
		log.Printf("[CreatePhoto - Repo] err: %s\n", err.Error())
		return photo, errs.NewInternalServerError("something went wrong")
	}

	return photo, nil
}

func (p *PhotoRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, photo entity.Photo) (entity.Photo, error) {
	sqlQuery := `UPDATE photos SET title=$1, caption=$2, photo_url=$3, updated_at=$4 WHERE id=$5 RETURNING updated_at, user_id`

	err := db.QueryRowContext(ctx, sqlQuery, photo.Title, photo.Caption, photo.PhotoUrl, time.Now(), photo.Id).Scan(&photo.UpdatedAt, &photo.UserId)

	if err != nil {
		log.Printf("[UpdatePhoto - Repo] err:%s\n", err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			return photo, errs.NewNotFoundError(fmt.Sprintf("photo with id: %v not found", photo.Id))
		}

		return photo, errs.NewInternalServerError("something went wrong")
	}

	return photo, nil
}

func (p *PhotoRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, photoId int) error {
	sqlQuery := `DELETE FROM photos WHERE id=$1 RETURNING id`

	var affectedRow int

	err := db.QueryRowContext(ctx, sqlQuery, photoId).Scan(&affectedRow)

	if err != nil {
		log.Printf("[DeletePhoto - Repo] err:%s\n", err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundError(fmt.Sprintf("photo with id: %v not found", photoId))
		}

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
