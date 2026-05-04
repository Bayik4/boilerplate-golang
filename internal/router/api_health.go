package router

import (
	"github.com/bayik4/boilerplate-golang/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerHealthRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/health")
	handler := handler.NewHandlerHealth()

	r.GET("/", handler.CheckHealth)
}