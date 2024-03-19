package service

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/repository"
)

type SocialMediaServiceImpl struct {
	SocialMediaRepository repository.SocialMediaRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewSocialMediaServiceImpl(socialMediaRepository repository.SocialMediaRepository, db *sql.DB, validate *validator.Validate) SocialMediaService {
	return &SocialMediaServiceImpl{
		SocialMediaRepository: socialMediaRepository,
		DB:                    db,
		Validate:              validate,
	}
}

func (s *SocialMediaServiceImpl) FindAll(ctx *gin.Context) (*[]dto.SocialMediaResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaServiceImpl) Create(ctx *gin.Context, socialMediaDto *dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaServiceImpl) Update(ctx *gin.Context, socialMediaDto *dto.UpdateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SocialMediaServiceImpl) Delete(ctx *gin.Context, socialMediaId int) error {
	panic("not implemented") // TODO: Implement
}
