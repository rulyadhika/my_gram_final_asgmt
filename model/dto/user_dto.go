package dto

import (
	"encoding/json"
	"time"
)

type NewUserRequest struct {
	Username string      `json:"username" validate:"required"`
	Email    string      `json:"email" validate:"required,email"`
	Password string      `json:"password" validate:"required"`
	Age      json.Number `json:"age" validate:"required,number"`
}

type NewUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

type UpdateUserRequest struct {
	Id       uint   `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      uint   `json:"age" validate:"required"`
}

type UpdateUserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
