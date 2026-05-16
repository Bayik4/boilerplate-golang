package handler

import (
	"net/http"

	"github.com/bayik4/boilerplate-golang/internal/service"
	"github.com/bayik4/boilerplate-golang/internal/utility"
	"github.com/gin-gonic/gin"
)

type HandlerHealth struct {
	healthService *service.HealthService
}

func newHandlerHealth(service *service.Service) *HandlerHealth {
	return &HandlerHealth{
		healthService: service.HealthService,
	}
}

func (h *HandlerHealth) CheckHealth(ctx *gin.Context) {
	res, _ := h.healthService.GetHealth(ctx)
	utility.ResponseWithoutData(ctx, res, http.StatusOK)
}