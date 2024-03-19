package service

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (u *UserServiceImpl) Register(ctx *gin.Context, userDto *dto.NewUserRequest) (*dto.NewUserResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserServiceImpl) Login(ctx *gin.Context, userDto *dto.LoginRequest) (*dto.LoginResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserServiceImpl) Update(ctx *gin.Context, userDto *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserServiceImpl) Delete(ctx *gin.Context, userId int) error {
	panic("not implemented") // TODO: Implement
}
