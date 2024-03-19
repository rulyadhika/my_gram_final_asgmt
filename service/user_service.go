package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
)

type UserService interface {
	Register(ctx *gin.Context, userDto *dto.NewUserRequest) (*dto.NewUserResponse, error)
	Login(ctx *gin.Context, userDto *dto.LoginRequest) (*dto.LoginResponse, error)
	Update(ctx *gin.Context, userDto *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
	Delete(ctx *gin.Context, userId int) error
}
