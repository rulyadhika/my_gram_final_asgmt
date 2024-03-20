package entity

import (
	"time"

	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
)

type SocialMedia struct {
	Id             uint
	Name           string
	SocialMediaUrl string
	UserId         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (s *SocialMedia) ToNewSocialMediaResponse() *dto.NewSocialMediaResponse {
	return &dto.NewSocialMediaResponse{
		Id:             s.Id,
		Name:           s.Name,
		SocialMediaUrl: s.SocialMediaUrl,
		UserId:         s.UserId,
		CreatedAt:      s.CreatedAt,
	}
}
