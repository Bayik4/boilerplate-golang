package router

import (
	"time"

	"github.com/bayik4/boilerplate-golang/internal/handler"
	"github.com/bayik4/boilerplate-golang/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(handler *handler.Handler, log *zap.Logger) *gin.Engine {
	r := gin.New()

	r.Use(
		middleware.ZapLogger(log),
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
		gin.Recovery(),
	)

	registerRoutes(r, handler)

	return r
}
