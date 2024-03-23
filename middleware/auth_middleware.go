package middleware

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
	UserRepository        repository.UserRepository
	DB                    *sql.DB
}

func NewAuthMiddlewareImpl(sr repository.SocialMediaRepository, pr repository.PhotoRepository, cr repository.CommentRepository, ur repository.UserRepository, db *sql.DB) AuthMiddleware {
	return &AuthMiddlewareImpl{
		SocialMediaRepository: sr,
		PhotoRepository:       pr,
		CommentRepository:     cr,
		UserRepository:        ur,
		DB:                    db,
	}
}

func (a *AuthMiddlewareImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")
		hasPrefixBearer := strings.HasPrefix(authorizationHeader, "Bearer")

		if !hasPrefixBearer {
			handleError(ctx, http.StatusUnauthorized, errs.NewUnauthorizedError("invalid token"))
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")

		if len(bearerToken) != 2 {
			handleError(ctx, http.StatusUnauthorized, errs.NewUnauthorizedError("invalid token"))
			return
		}

		token := bearerToken[1]

		user := entity.User{}

		err := user.ParseToken(token)

		if err != nil {
			handleError(ctx, http.StatusUnauthorized, err)
			return
		}

		_, err = a.UserRepository.GetUserById(ctx, a.DB, int(user.Id))
		if err != nil {
			handleError(ctx, http.StatusUnauthorized, errs.NewUnauthorizedError("invalid token"))
			return
		}

		ctx.Set("userData", user)

		ctx.Next()
	}
}

func (a *AuthMiddlewareImpl) AuthorizationSocialMedia() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.MustGet("userData").(entity.User)

		if !ok {
			log.Printf("[AuthorizationSocialMedia - Middleware] err:%s\n", "failed type casting to 'entity.user'")
			handleError(ctx, http.StatusInternalServerError, errs.NewInternalServerError("something went wrong"))
			return
		}

		socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))

		if err != nil {
			log.Printf("[AuthorizationSocialMedia - Middleware] err:%s\n", err.Error())
			handleError(ctx, http.StatusInternalServerError, errs.NewInternalServerError("something went wrong"))
			return
		}

		socialMedia, err := a.SocialMediaRepository.FindById(ctx, a.DB, socialMediaId)
		if err != nil {
			log.Printf("[AuthorizationSocialMedia - Middleware] err:%s\n", err.Error())
			handleError(ctx, http.StatusInternalServerError, errs.NewInternalServerError("something went wrong"))
			return
		}

		if socialMedia.UserId != user.Id {
			handleError(ctx, http.StatusForbidden, errs.NewForbiddenError("you are not authorized to access/modify this data"))
			return
		}

		ctx.Next()
	}
}

func (a *AuthMiddlewareImpl) AuthorizationPhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.MustGet("userData").(entity.User)

		if !ok {
			log.Printf("[AuthorizationPhoto - Middleware] err:%s\n", "failed type casting to 'entity.user'")
			handleError(ctx, http.StatusInternalServerError, errs.NewInternalServerError("something went wrong"))
			return
		}

		photoId, err := strconv.Atoi(ctx.Param("photoId"))

		if err != nil {
			log.Printf("[AuthorizationPhoto - Middleware] err:%s\n", err.Error())
			handleError(ctx, http.StatusInternalServerError, errs.NewInternalServerError("something went wrong"))
			return
		}

		photo, err := a.PhotoRepository.FindById(ctx, a.DB, photoId)
		if err != nil {
			log.Printf("[AuthorizationPhoto - Middleware] err:%s\n", err.Error())
			handleError(ctx, http.StatusInternalServerError, errs.NewInternalServerError("something went wrong"))
			return
		}

		if photo.UserId != user.Id {
			handleError(ctx, http.StatusForbidden, errs.NewForbiddenError("you are not authorized to access/modify this data"))
			return
		}

		ctx.Next()
	}
}

func (a *AuthMiddlewareImpl) AuthorizationComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.MustGet("userData").(entity.User)

		if !ok {
			log.Printf("[AuthorizationComment - Middleware] err:%s\n", "failed type casting to 'entity.user'")
			handleError(ctx, http.StatusInternalServerError, errs.NewInternalServerError("something went wrong"))
			return
		}

		commentId, err := strconv.Atoi(ctx.Param("commentId"))

		if err != nil {
			log.Printf("[AuthorizationComment - Middleware] err:%s\n", err.Error())
			handleError(ctx, http.StatusInternalServerError, errs.NewInternalServerError("something went wrong"))
			return
		}

		photo, err := a.CommentRepository.FindById(ctx, a.DB, commentId)
		if err != nil {
			log.Printf("[AuthorizationComment - Middleware] err:%s\n", err.Error())
			handleError(ctx, http.StatusInternalServerError, errs.NewInternalServerError("something went wrong"))
			return
		}

		if photo.UserId != user.Id {
			handleError(ctx, http.StatusForbidden, errs.NewForbiddenError("you are not authorized to access/modify this data"))
			return
		}

		ctx.Next()
	}
}
