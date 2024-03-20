package dto

import "time"

type NewSocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
	UserId         uint   `json:"user_id" validate:"required"`
}

type NewSocialMediaResponse struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaResponse struct {
	Id             uint         `json:"id"`
	Name           string       `json:"name"`
	SocialMediaUrl string       `json:"social_media_url"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	User           UserResponse `json:"user"`
}

type UpdateSocialMediaRequest struct {
	Id             uint   `json:"id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

type UpdateSocialMediaResponse struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}
