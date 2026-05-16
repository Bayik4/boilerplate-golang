package router

import (
	"github.com/bayik4/boilerplate-golang/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, handler *handler.Handler) {
	api := r.Group("/api/v1")
	
	registerHealthRoutes(api, handler.HandlerHealth)
}