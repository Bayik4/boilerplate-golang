package cmd

import (
	"context"
	"fmt"

	"github.com/bayik4/boilerplate-golang/boot"
	"github.com/bayik4/boilerplate-golang/internal/router"
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

		r := router.NewRouter(log)

		addr := fmt.Sprintf(":%s", cfg.App.Port)

		log.Info("starting http server",
			zap.String("port", cfg.App.Port),
		)

		if err := r.Run(addr); err != nil {
			log.Fatal("failed to start server", zap.Error(err))
		}
	},
}
