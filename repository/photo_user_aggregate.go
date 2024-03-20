package repository

import "github.com/rulyadhika/my_gram_final_asgmt/model/entity"

type PhotoUser struct {
	entity.Photo
	entity.User
}
