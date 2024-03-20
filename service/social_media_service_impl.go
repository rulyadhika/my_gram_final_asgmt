package service

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/helper"
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
	result, err := s.SocialMediaRepository.FindAll(ctx, s.DB)

	if err != nil {
		return &[]dto.SocialMediaResponse{}, err
	}

	return helper.ToSocialMediasResponse(result), nil
}

func (s *SocialMediaServiceImpl) Create(ctx *gin.Context, socialMediaDto *dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, error) {
	validationErr := s.Validate.Struct(socialMediaDto)

	if validationErr != nil {
		return &dto.NewSocialMediaResponse{}, validationErr
	}

	socialMediaData := entity.SocialMedia{
		Name:           socialMediaDto.Name,
		SocialMediaUrl: socialMediaDto.SocialMediaUrl,
		UserId:         socialMediaDto.UserId,
	}

	result, err := s.SocialMediaRepository.Create(ctx, s.DB, socialMediaData)

	if err != nil {
		return &dto.NewSocialMediaResponse{}, err
	}

	return helper.ToNewSocialMediaResponse(result), nil
}

func (s *SocialMediaServiceImpl) Update(ctx *gin.Context, socialMediaDto *dto.UpdateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, error) {
	validationErr := s.Validate.Struct(socialMediaDto)

	if validationErr != nil {
		return &dto.UpdateSocialMediaResponse{}, validationErr
	}

	// TODO change user id get from context
	userId := 1

	socialMedia := entity.SocialMedia{
		Id:             socialMediaDto.Id,
		Name:           socialMediaDto.Name,
		SocialMediaUrl: socialMediaDto.SocialMediaUrl,
		UserId:         uint(userId),
	}

	result, err := s.SocialMediaRepository.Update(ctx, s.DB, socialMedia)

	if err != nil {
		return &dto.UpdateSocialMediaResponse{}, err
	}

	return helper.ToUpdateSocialMediaResponse(result), nil
}

func (s *SocialMediaServiceImpl) Delete(ctx *gin.Context, socialMediaId int) error {
	err := s.SocialMediaRepository.Delete(ctx, s.DB, socialMediaId)

	if err != nil {
		return err
	}

	return nil
}
