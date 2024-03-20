package entity

import (
	"time"
)

type SocialMedia struct {
	Id             uint
	Name           string
	SocialMediaUrl string
	UserId         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
