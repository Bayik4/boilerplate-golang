package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/bayik4/boilerplate-golang/boot"
	"github.com/bayik4/boilerplate-golang/internal/handler"
	"github.com/bayik4/boilerplate-golang/internal/repository"
	"github.com/bayik4/boilerplate-golang/internal/repository/util"
	"github.com/bayik4/boilerplate-golang/internal/router"
	"github.com/bayik4/boilerplate-golang/internal/service"
	"github.com/bayik4/boilerplate-golang/internal/telemetry"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var restCommand = &cobra.Command{
	Use:   "serve",
	Short: "Run API Service",
	Run: func(_ *cobra.Command, args []string) {

		ctx := context.Background()
		vpr := viper.New()

		logConfig := boot.LoadLoggerConfig(vpr)

		log, err := boot.LoggerInit(logConfig)
		if err != nil {
			log.Fatal("failed to load logger", zap.Error(err))
		}
		defer func() { _ = log.Sync() }()
		log.Info("Running application")

		shutdown := func(context.Context) error { return nil }

		cfg := boot.LoadConfig(ctx, vpr, log)
		cfg.Logger = logConfig

		if cfg.Otel.Enabled {
			shutdown, err = telemetry.Init(ctx, cfg.Otel)
			if err != nil {
				log.Fatal("failed to initialize telemetry", zap.Error(err))
			}
		}
		defer func() {
			shutdownCtx, cancel := context.WithTimeout(
				context.Background(), 5*time.Second,
			)
			defer cancel()
			if err := shutdown(shutdownCtx); err != nil {
				log.Error("failed to shutdown telemetry", zap.Error(err))
			}
		}()

		dbexec := util.NewDBExecutor(log)
		repo := repository.New(cfg.Database.Pgsql.DBInit, *dbexec, log)
		serv := service.New(repo, log)
		handler := handler.New(serv, log)
		r := router.NewRouter(handler, log, *cfg)

		addr := fmt.Sprintf(":%s", cfg.App.Port)

		log.Info("starting http server",
			zap.String("port", cfg.App.Port),
		)

		if err := r.Run(addr); err != nil {
			log.Fatal("failed to start server", zap.Error(err))
		}
	},
}
