package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
	"github.com/rulyadhika/my_gram_final_asgmt/service"
)

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandlerImpl(userService service.UserService) UserHandler {
	return &UserHandlerImpl{UserService: userService}
}

func (u *UserHandlerImpl) Register(ctx *gin.Context) {
	userDto := &dto.NewUserRequest{}

	err := ctx.ShouldBindJSON(userDto)

	if err != nil {
		log.Printf("[RegisterUser - Handler] err:%s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	result, err := u.UserService.Register(ctx, userDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusCreated),
		Code:    http.StatusCreated,
		Message: "successfully register new user",
		Data:    result,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (u *UserHandlerImpl) Login(ctx *gin.Context) {
	userDto := &dto.LoginRequest{}

	err := ctx.ShouldBindJSON(userDto)

	if err != nil {
		log.Printf("[LoginUser - Handler] err:%s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	result, err := u.UserService.Login(ctx, userDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "successfully login",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (u *UserHandlerImpl) Update(ctx *gin.Context) {
	userDto := &dto.UpdateUserRequest{}

	err := ctx.ShouldBindJSON(userDto)

	if err != nil {
		log.Printf("[UpdateUser - Handler] err:%s\n", err.Error())
		ctx.Error(errs.NewUnprocessableEntityError("invalid json request body"))
		return
	}

	user, ok := ctx.MustGet("userData").(entity.User)
	if !ok {
		log.Printf("[UpdateUser - Handler] err: failed type casting to 'entity.user'\n")
		ctx.Error(errs.NewInternalServerError("something went wrong"))
		return
	}

	userDto.Id = user.Id

	result, err := u.UserService.Update(ctx, userDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "successfully update user",
		Data:    result,
	}

	ctx.JSON(http.StatusOK, response)
}

func (u *UserHandlerImpl) Delete(ctx *gin.Context) {
	user, ok := ctx.MustGet("userData").(entity.User)
	if !ok {
		log.Printf("[DeleteUser - handler] err: failed type casting to 'entity.user'\n")
		ctx.Error(errs.NewInternalServerError("something went wrong"))
		return
	}

	userId := user.Id

	err := u.UserService.Delete(ctx, int(userId))

	if err != nil {
		ctx.Error(err)
		return
	}

	response := &dto.WebResponse{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "your account has been successfully deleted",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, response)
}
