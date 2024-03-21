package repository

import "github.com/rulyadhika/my_gram_final_asgmt/model/entity"

type CommentPhotoUser struct {
	entity.Comment
	entity.User
	PhotoUser PhotoUser
}
