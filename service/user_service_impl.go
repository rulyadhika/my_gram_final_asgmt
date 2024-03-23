package service

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/helper"
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
	validationErr := u.Validate.Struct(userDto)

	if validationErr != nil {
		return &dto.NewUserResponse{}, validationErr
	}

	age, err := userDto.Age.Float64()

	if err != nil {
		return &dto.NewUserResponse{}, errs.NewUnprocessableEntityError("age must be a number")
	}

	user := entity.User{
		Username: userDto.Username,
		Email:    userDto.Email,
		Age:      uint(age),
		Password: userDto.Password,
	}

	// check email and username should be unique
	err = u.UserRepository.CheckEmailAndUsernameUnique(ctx, u.DB, user)

	if err != nil {
		return &dto.NewUserResponse{}, err
	}

	// hash password
	err = user.HashPassword()
	if err != nil {
		return &dto.NewUserResponse{}, err
	}

	result, err := u.UserRepository.Register(ctx, u.DB, user)

	if err != nil {
		return &dto.NewUserResponse{}, err
	}

	return helper.ToNewUserResponse(result), nil
}

func (u *UserServiceImpl) Login(ctx *gin.Context, userDto *dto.LoginRequest) (*dto.LoginResponse, error) {
	validationErr := u.Validate.Struct(userDto)

	if validationErr != nil {
		return &dto.LoginResponse{}, validationErr
	}

	user := &entity.User{
		Email: userDto.Email,
	}

	// check if email exists
	err := u.UserRepository.GetUserByEmail(ctx, u.DB, user)

	if err != nil {
		switch err.(type) {
		case *errs.NotFoundError:
			return &dto.LoginResponse{}, errs.NewBadRequestError("invalid email/password")
		default:
			return &dto.LoginResponse{}, err
		}
	}

	// validate password
	passwordValid := user.ValidatePassword(userDto.Password)

	if !passwordValid {
		return &dto.LoginResponse{}, errs.NewBadRequestError("invalid email/password")
	}

	// generate jwt token
	token, err := user.GenerateToken()

	if err != nil {
		return &dto.LoginResponse{}, err
	}

	return helper.ToLoginResponse(token.(string)), nil
}

func (u *UserServiceImpl) Update(ctx *gin.Context, userDto *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	validationErr := u.Validate.Struct(userDto)

	if validationErr != nil {
		return &dto.UpdateUserResponse{}, validationErr
	}

	age, err := userDto.Age.Float64()

	if err != nil {
		return &dto.UpdateUserResponse{}, errs.NewUnprocessableEntityError("age must be a number")
	}

	user := entity.User{
		Id:       userDto.Id,
		Username: userDto.Username,
		Email:    userDto.Email,
		Age:      uint(age),
	}

	err = u.UserRepository.CheckEmailAndUsernameUnique(ctx, u.DB, user)
	if err != nil {
		return &dto.UpdateUserResponse{}, err
	}

	result, err := u.UserRepository.Update(ctx, u.DB, user)

	if err != nil {
		return &dto.UpdateUserResponse{}, err
	}

	return helper.ToUpdateUserResponse(result), nil
}

func (u *UserServiceImpl) Delete(ctx *gin.Context, userId int) error {
	err := u.UserRepository.Delete(ctx, u.DB, userId)

	if err != nil {
		return err
	}

	return nil
}
