package router

import "github.com/gin-gonic/gin"

func registerRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	
	registerHealthRoutes(api)
}