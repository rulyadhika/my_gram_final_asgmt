package entity

import "time"

type Comment struct {
	Id        uint
	UserId    uint
	PhotoId   uint
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
