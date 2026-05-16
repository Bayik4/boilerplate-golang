package cmd

import (
	"context"
	"fmt"

	"github.com/bayik4/boilerplate-golang/boot"
	"github.com/bayik4/boilerplate-golang/internal/handler"
	"github.com/bayik4/boilerplate-golang/internal/repository"
	"github.com/bayik4/boilerplate-golang/internal/repository/util"
	"github.com/bayik4/boilerplate-golang/internal/router"
	"github.com/bayik4/boilerplate-golang/internal/service"
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
		
		log := boot.LoggerInit()
		log.Info("Running application")

		cfg := boot.LoadConfig(ctx, vpr, log)

		dbexec := util.NewDBExecutor(log)
		repo := repository.New(cfg.Database.Pgsql.DBInit, *dbexec)
		serv := service.New(repo)
		handler := handler.New(serv)
		r := router.NewRouter(handler, log)

		addr := fmt.Sprintf(":%s", cfg.App.Port)

		log.Info("starting http server",
			zap.String("port", cfg.App.Port),
		)

		if err := r.Run(addr); err != nil {
			log.Fatal("failed to start server", zap.Error(err))
		}
	},
}
