package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
)

type SocialMediaRepositoryImpl struct{}

func (s *SocialMediaRepositoryImpl) FindAll(ctx *gin.Context, db *sql.DB) (*[]entity.SocialMedia, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaRepositoryImpl) Create(ctx *gin.Context, db *sql.DB, socialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, socialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, socialMediaId int) error {
	panic("not implemented") // TODO: Implement
}
