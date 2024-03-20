package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
			case validator.ValidationErrors:
				handleValidationError(ctx, err.Err.(validator.ValidationErrors))
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

func handleValidationError(ctx *gin.Context, errs validator.ValidationErrors) {
	var errors string

	for index, err := range errs {
		errors += fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag())

		if index != len(errs)-1 {
			errors += ","
		}
	}

	webResponse := &dto.WebResponse{
		Status:  http.StatusText(http.StatusBadRequest),
		Code:    http.StatusBadRequest,
		Message: errors,
		Data:    nil,
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)
}
