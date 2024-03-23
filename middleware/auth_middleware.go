package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
	"github.com/rulyadhika/my_gram_final_asgmt/repository"
)

type AuthMiddleware interface {
	Authentication() gin.HandlerFunc
	AuthorizationSocialMedia() gin.HandlerFunc
	AuthorizationPhoto() gin.HandlerFunc
	AuthorizationComment() gin.HandlerFunc
}

type AuthMiddlewareImpl struct {
	SocialMediaRepository repository.SocialMediaRepository
	PhotoRepository       repository.PhotoRepository
	CommentRepository     repository.CommentRepository
}

func NewAuthMiddlewareImpl(sr repository.SocialMediaRepository, pr repository.PhotoRepository, cr repository.CommentRepository) AuthMiddleware {
	return &AuthMiddlewareImpl{
		SocialMediaRepository: sr,
		PhotoRepository:       pr,
		CommentRepository:     cr,
	}
}

func (a *AuthMiddlewareImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")
		hasPrefixBearer := strings.HasPrefix(authorizationHeader, "Bearer")

		if !hasPrefixBearer {
			HandleUnauthorizedError(ctx, errs.NewUnauthorizedError("invalid token"))
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")

		if len(bearerToken) != 2 {
			HandleUnauthorizedError(ctx, errs.NewUnauthorizedError("invalid token"))
			return
		}

		token := bearerToken[1]

		user := entity.User{}

		err := user.ParseToken(token)

		if err != nil {
			HandleUnauthorizedError(ctx, err)
			return
		}

		// TODO
		// get user by id
		// if nil handle authorization error
		// ctx.Set("userData", user)

		ctx.Next()
	}
}

func (a *AuthMiddlewareImpl) AuthorizationSocialMedia() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func (a *AuthMiddlewareImpl) AuthorizationPhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func (a *AuthMiddlewareImpl) AuthorizationComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func HandleUnauthorizedError(ctx *gin.Context, err error) {
	webResponse := &dto.WebResponse{
		Status:  http.StatusText(http.StatusUnauthorized),
		Code:    http.StatusUnauthorized,
		Message: err.Error(),
		Data:    nil,
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
}
