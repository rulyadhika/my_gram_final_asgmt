package dto

import (
	"encoding/json"
	"time"
)

type NewCommentRequest struct {
	UserId  uint        `json:"user_id" validate:"required"`
	PhotoId json.Number `json:"photo_id" validate:"required"`
	Message string      `json:"message" validate:"required"`
}

type NewCommentResponse struct {
	Id        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   uint      `json:"photo_id"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateCommentRequest struct {
	Id      uint   `json:"id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type UpdateCommentResponse struct {
	Id        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   uint      `json:"photo_id"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentResponse struct {
	Id        uint          `json:"id"`
	Message   string        `json:"message"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      UserResponse  `json:"user"`
	Photo     PhotoResponse `json:"photo"`
}
