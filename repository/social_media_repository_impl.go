package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
)

type SocialMediaRepositoryImpl struct{}

func NewSocialMediaRepositoryImpl() SocialMediaRepository {
	return &SocialMediaRepositoryImpl{}
}

func (s *SocialMediaRepositoryImpl) FindAll(ctx *gin.Context, db *sql.DB) ([]entity.SocialMedia, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaRepositoryImpl) Create(ctx *gin.Context, db *sql.DB, socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	sqlQuery := "INSERT INTO social_medias(name, social_media_url, user_id) VALUES($1, $2, $3) RETURNING id"

	err := db.QueryRowContext(ctx, sqlQuery, socialMedia.Name, socialMedia.SocialMediaUrl, socialMedia.UserId).Scan(socialMedia.Id)

	if err != nil {
		return socialMedia, errs.NewInternalServerError("something went wrong")
	}

	return socialMedia, nil
}

func (s *SocialMediaRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, socialMediaId int) error {
	panic("not implemented") // TODO: Implement
}
