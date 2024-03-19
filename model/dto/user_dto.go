package dto

import "time"

type NewUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min:6"`
	Age      uint   `json:"age" validate:"required,min:8"`
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
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username uint   `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password uint   `json:"password" validate:"required,min:6"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
