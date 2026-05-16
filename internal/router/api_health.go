package router

import (
	"github.com/bayik4/boilerplate-golang/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerHealthRoutes(rg *gin.RouterGroup, handler *handler.HandlerHealth) {
	r := rg.Group("/health")

	r.GET("/", handler.CheckHealth)
}