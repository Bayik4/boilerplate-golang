package handler

import (
	"net/http"

	"github.com/bayik4/boilerplate-golang/internal/utility"
	"github.com/gin-gonic/gin"
)

type HandlerHealth struct {}

func NewHandlerHealth() *HandlerHealth {
	return &HandlerHealth{}
}

func (_ *HandlerHealth) CheckHealth(ctx *gin.Context) {
	utility.ResponseWithoutData(ctx, "ok", http.StatusOK)
}