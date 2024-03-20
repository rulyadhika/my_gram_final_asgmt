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

type SocialMediaRepositoryImpl struct{}

func NewSocialMediaRepositoryImpl() SocialMediaRepository {
	return &SocialMediaRepositoryImpl{}
}

func (s *SocialMediaRepositoryImpl) FindAll(ctx *gin.Context, db *sql.DB) ([]SocialMediaUser, error) {
	sqlQuery := `SELECT social_medias.id, social_medias.name, social_medias.social_media_url,
	social_medias.user_id, social_medias.created_at, social_medias.updated_at, users.id, users.email, users.username
	FROM social_medias JOIN users ON social_medias.user_id = users.id`

	socialMediasUser := []SocialMediaUser{}

	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		log.Printf("[FindAllSocialMedia - Repo] err: %s \n", err.Error())
		return socialMediasUser, errs.NewInternalServerError("something went wrong")
	}
	defer rows.Close()

	for rows.Next() {
		socialMediaUser := SocialMediaUser{}

		err := rows.Scan(&socialMediaUser.SocialMedia.Id, &socialMediaUser.Name, &socialMediaUser.SocialMediaUrl, &socialMediaUser.UserId, &socialMediaUser.SocialMedia.CreatedAt, &socialMediaUser.SocialMedia.UpdatedAt, &socialMediaUser.User.Id, &socialMediaUser.Email, &socialMediaUser.Username)

		if err != nil {
			log.Printf("[FindAllSocialMedia - Repo] err: %s \n", err.Error())
			return socialMediasUser, errs.NewInternalServerError("something went wrong")
		}

		socialMediasUser = append(socialMediasUser, socialMediaUser)
	}

	return socialMediasUser, nil
}

func (s *SocialMediaRepositoryImpl) Create(ctx *gin.Context, db *sql.DB, socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	sqlQuery := "INSERT INTO social_medias(name, social_media_url, user_id) VALUES($1, $2, $3) RETURNING id, created_at"

	err := db.QueryRowContext(ctx, sqlQuery, socialMedia.Name, socialMedia.SocialMediaUrl, socialMedia.UserId).Scan(&socialMedia.Id, &socialMedia.CreatedAt)

	if err != nil {
		log.Printf("[CreateSocialMedia - Repo] err: %s \n", err.Error())
		return socialMedia, errs.NewInternalServerError("something went wrong")
	}

	return socialMedia, nil
}

func (s *SocialMediaRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	sqlQuery := `UPDATE social_medias SET name=$1, social_media_url=$2, updated_at=$3 WHERE id=$4 RETURNING updated_at`

	err := db.QueryRowContext(ctx, sqlQuery, socialMedia.Name, socialMedia.SocialMediaUrl, time.Now(), socialMedia.Id).Scan(&socialMedia.UpdatedAt)

	if err != nil {
		log.Printf("[UpdateSocialMedia - Repo] err: %s \n", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return socialMedia, errs.NewNotFoundError(fmt.Sprintf("social media with id:%v not found", socialMedia.Id))
		}

		return socialMedia, errs.NewInternalServerError("something went wrong")
	}

	return socialMedia, nil
}

func (s *SocialMediaRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, socialMediaId int) error {
	sqlQuery := `DELETE FROM social_medias WHERE id=$1 RETURNING id`

	var affectedRow int

	err := db.QueryRowContext(ctx, sqlQuery, socialMediaId).Scan(&affectedRow)

	if err != nil {
		log.Printf("[DeleteSocialMedia - Repo] err: %s \n", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundError(fmt.Sprintf("social media with id:%v not found", socialMediaId))
		}

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
