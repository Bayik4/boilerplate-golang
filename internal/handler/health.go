package handler

import (
	"net/http"

	"github.com/bayik4/boilerplate-golang/internal/service"
	"github.com/bayik4/boilerplate-golang/internal/utility"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HandlerHealth struct {
	healthService *service.HealthService
	logger        *zap.Logger
}

func newHandlerHealth(service *service.Service, logger *zap.Logger) *HandlerHealth {
	return &HandlerHealth{
		healthService: service.HealthService,
		logger:        logger,
	}
}

func (h *HandlerHealth) CheckHealth(ctx *gin.Context) {
	res, _ := h.healthService.GetHealth(ctx)
	h.logger.Info("Application is healty")
	utility.ResponseWithoutData(ctx, res, http.StatusOK)
}
