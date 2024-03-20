package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/dto"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()

		if err != nil {
			switch err.Err.(type) {
			case *errs.UnprocessableEntityError:
				handleError(ctx, http.StatusUnprocessableEntity, err)
			case *errs.BadRequestError:
				handleError(ctx, http.StatusBadRequest, err)
			case *errs.NotFoundError:
				handleError(ctx, http.StatusNotFound, err)
			default:
				handleError(ctx, http.StatusInternalServerError, err)
			}
		}
	}
}

func handleError(ctx *gin.Context, httpStatusCode int, err error) {
	webResponse := &dto.WebResponse{
		Status:  http.StatusText(httpStatusCode),
		Code:    httpStatusCode,
		Message: err.Error(),
		Data:    nil,
	}

	ctx.AbortWithStatusJSON(httpStatusCode, webResponse)
}
