package helper

import (
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
)

func ToNewPhotoResponse(p entity.Photo) *dto.NewPhotoResponse {
	return &dto.NewPhotoResponse{
		Id:        p.Id,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoUrl:  p.PhotoUrl,
		UserId:    p.UserId,
		CreatedAt: p.CreatedAt,
	}
}
