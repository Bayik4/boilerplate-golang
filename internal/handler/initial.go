package handler

import "github.com/bayik4/boilerplate-golang/internal/service"

type Handler struct {
	HandlerHealth *HandlerHealth
}

func New(service *service.Service) *Handler {
	return &Handler{
		HandlerHealth: newHandlerHealth(service),
	}
}