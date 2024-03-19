package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/service"
)

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandlerImpl(userService service.UserService) UserHandler {
	return &UserHandlerImpl{UserService: userService}
}

func (u *UserHandlerImpl) Register(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (u *UserHandlerImpl) Login(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (u *UserHandlerImpl) Update(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (u *UserHandlerImpl) Delete(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}
