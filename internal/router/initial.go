package router

import (
	"time"

	"github.com/bayik4/boilerplate-golang/internal/handler"
	"github.com/bayik4/boilerplate-golang/internal/middleware"
	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

func NewRouter(handler *handler.Handler, log *zap.Logger, cfg model.ConfigModel) *gin.Engine {
	r := gin.New()

	gin.SetMode(gin.ReleaseMode)
	if cfg.App.Env == "development" {
		gin.SetMode(gin.DebugMode)
	}
	r.Use(
		middleware.RequestID(),
		otelgin.Middleware(cfg.App.Name),
		middleware.RequestLogger(log),
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
