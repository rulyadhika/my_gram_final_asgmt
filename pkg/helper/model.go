package helper

import (
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/repository"
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

func ToPhotoResponse(p []repository.PhotoUser) *[]dto.PhotoResponse {
	photosResponse := []dto.PhotoResponse{}

	for _, item := range p {
		user := dto.UserResponse{
			Id:       item.User.Id,
			Email:    item.Email,
			Username: item.Username,
		}

		photoResponse := dto.PhotoResponse{
			Id:        item.Photo.Id,
			Title:     item.Title,
			Caption:   item.Caption,
			PhotoUrl:  item.PhotoUrl,
			CreatedAt: item.Photo.CreatedAt,
			UpdatedAt: item.Photo.UpdatedAt,
			User:      user,
		}

		photosResponse = append(photosResponse, photoResponse)
	}

	return &photosResponse
}

func ToPhotoUpdateResponse(p entity.Photo) *dto.UpdatePhotoResponse {
	return &dto.UpdatePhotoResponse{
		Id:        p.Id,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoUrl:  p.PhotoUrl,
		UserId:    p.UserId,
		UpdatedAt: p.UpdatedAt,
	}
}

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

func ToNewCommentResponse(c entity.Comment) *dto.NewCommentResponse {
	return &dto.NewCommentResponse{
		Id:        c.Id,
		Message:   c.Message,
		PhotoId:   c.PhotoId,
		UserId:    c.UserId,
		CreatedAt: c.CreatedAt,
	}
}

func ToCommentsResponse(c []repository.CommentPhotoUser) *[]dto.CommentResponse {
	comments := []dto.CommentResponse{}

	for _, item := range c {
		userComment := dto.UserResponse{
			Id:       item.User.Id,
			Email:    item.Email,
			Username: item.Username,
		}

		userPhoto := dto.UserResponse{
			Id:       item.PhotoUser.User.Id,
			Email:    item.PhotoUser.User.Email,
			Username: item.PhotoUser.User.Username,
		}

		photo := dto.PhotoResponse{
			Id:        item.PhotoUser.Photo.Id,
			Title:     item.PhotoUser.Photo.Title,
			Caption:   item.PhotoUser.Photo.Caption,
			PhotoUrl:  item.PhotoUser.Photo.PhotoUrl,
			CreatedAt: item.PhotoUser.Photo.CreatedAt,
			UpdatedAt: item.PhotoUser.Photo.UpdatedAt,
			User:      userPhoto,
		}

		comment := dto.CommentResponse{
			Id:        item.Comment.Id,
			Message:   item.Message,
			CreatedAt: item.Comment.CreatedAt,
			UpdatedAt: item.Comment.UpdatedAt,
			User:      userComment,
			Photo:     photo,
		}

		comments = append(comments, comment)
	}

	return &comments
}

func ToUpdateCommentResponse(c entity.Comment) *dto.UpdateCommentResponse {
	return &dto.UpdateCommentResponse{
		Id:        c.Id,
		Message:   c.Message,
		PhotoId:   c.PhotoId,
		UserId:    c.UserId,
		UpdatedAt: c.UpdatedAt,
	}
}

func ToNewUserResponse(u entity.User) *dto.NewUserResponse {
	return &dto.NewUserResponse{
		Username: u.Username,
		Email:    u.Email,
		Age:      u.Age,
	}
}

func ToLoginResponse(token string) *dto.LoginResponse {
	return &dto.LoginResponse{
		Token: token,
	}
}
