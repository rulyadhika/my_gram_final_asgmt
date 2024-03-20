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

type PhotoServiceImpl struct {
	PhotoRepository repository.PhotoRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewPhotoServiceImpl(photoRepository repository.PhotoRepository, db *sql.DB, validate *validator.Validate) PhotoService {
	return &PhotoServiceImpl{
		PhotoRepository: photoRepository,
		DB:              db,
		Validate:        validate,
	}
}

func (p *PhotoServiceImpl) FindAll(ctx *gin.Context) (*[]dto.PhotoResponse, error) {
	result, err := p.PhotoRepository.FindAll(ctx, p.DB)

	if err != nil {
		return &[]dto.PhotoResponse{}, err
	}

	return helper.ToPhotoResponse(result), nil

}

func (p *PhotoServiceImpl) Create(ctx *gin.Context, photoDto *dto.NewPhotoRequest) (*dto.NewPhotoResponse, error) {
	validationErr := p.Validate.Struct(photoDto)

	if validationErr != nil {
		return &dto.NewPhotoResponse{}, validationErr
	}

	photo := entity.Photo{
		Title:    photoDto.Title,
		Caption:  photoDto.Caption,
		PhotoUrl: photoDto.PhotoUrl,
		UserId:   photoDto.UserId,
	}

	result, err := p.PhotoRepository.Create(ctx, p.DB, photo)

	if err != nil {
		return &dto.NewPhotoResponse{}, err
	}

	return helper.ToNewPhotoResponse(result), nil
}

func (p *PhotoServiceImpl) Update(ctx *gin.Context, photoDto *dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, error) {
	validationErr := p.Validate.Struct(photoDto)

	if validationErr != nil {
		return &dto.UpdatePhotoResponse{}, validationErr
	}

	photo := entity.Photo{
		Id:       photoDto.Id,
		Title:    photoDto.Title,
		Caption:  photoDto.Caption,
		PhotoUrl: photoDto.PhotoUrl,
	}

	result, err := p.PhotoRepository.Update(ctx, p.DB, photo)

	if err != nil {
		return &dto.UpdatePhotoResponse{}, err
	}

	return helper.ToPhotoUpdateResponse(result), nil
}

func (p *PhotoServiceImpl) Delete(ctx *gin.Context, photoId int) error {
	err := p.PhotoRepository.Delete(ctx, p.DB, photoId)

	if err != nil {
		return err
	}

	return nil
}
