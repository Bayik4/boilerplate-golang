package utility

import (
	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/gin-gonic/gin"
)

func ResponseWithData[T any](ctx *gin.Context, data T, msg string, code int) {
	ctx.JSON(code, model.ResponseApi[T]{
		Status: true,
		Message: msg,
		Data: data,
	})
}

func ResponseWithoutData(ctx *gin.Context, msg string, code int) {
	ctx.JSON(code, model.ResponseApiWithoutData{
		Status: true,
		Message: msg,
	})
}

func ResponseError(ctx *gin.Context, err string, msg string, code int) {
	ctx.JSON(code, model.ResponseApiError{
		Status: true,
		Message: msg,
		Error: err,
	})
}