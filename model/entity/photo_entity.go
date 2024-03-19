package entity

import "time"

type Photo struct {
	Id        uint
	Title     string
	Caption   string
	PhotoUrl  string
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
