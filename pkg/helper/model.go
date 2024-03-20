package helper

import (
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/repository"
)

func ToNewSocialMediaResponse(s entity.SocialMedia) *dto.NewSocialMediaResponse {
	return &dto.NewSocialMediaResponse{
		Id:             s.Id,
		Name:           s.Name,
		SocialMediaUrl: s.SocialMediaUrl,
		UserId:         s.UserId,
		CreatedAt:      s.CreatedAt,
	}
}

func ToSocialMediasResponse(s []repository.SocialMediaUser) *[]dto.SocialMediaResponse {
	socialMedias := []dto.SocialMediaResponse{}

	for _, item := range s {
		user := dto.UserResponse{
			Id:       item.User.Id,
			Email:    item.Email,
			Username: item.Username,
		}

		socialMedia := dto.SocialMediaResponse{
			Id:             item.SocialMedia.Id,
			Name:           item.Name,
			SocialMediaUrl: item.SocialMediaUrl,
			CreatedAt:      item.SocialMedia.CreatedAt,
			UpdatedAt:      item.SocialMedia.UpdatedAt,
			User:           user,
		}

		socialMedias = append(socialMedias, socialMedia)
	}

	return &socialMedias
}

func ToUpdateSocialMediaResponse(s entity.SocialMedia) *dto.UpdateSocialMediaResponse {
	return &dto.UpdateSocialMediaResponse{
		Id:             s.Id,
		Name:           s.Name,
		SocialMediaUrl: s.SocialMediaUrl,
		UserId:         s.UserId,
		UpdatedAt:      s.UpdatedAt,
	}
}
