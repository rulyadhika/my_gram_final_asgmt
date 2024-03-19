package entity

import "time"

type User struct {
	Id        uint
	Username  string
	Email     string
	Password  string
	Age       uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
