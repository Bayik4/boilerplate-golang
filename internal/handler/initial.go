package handler

import (
	"github.com/bayik4/boilerplate-golang/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
	HandlerHealth *HandlerHealth
}

func New(service *service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		HandlerHealth: newHandlerHealth(service, logger),
	}
}
