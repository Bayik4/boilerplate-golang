package boot

import (
	"context"

	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func LoadConfig(ctx context.Context, vpr *viper.Viper, log *zap.Logger) *model.ConfigModel {
	vpr.SetConfigName("config")
	vpr.SetConfigType("yaml")
	vpr.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := vpr.ReadInConfig(); err != nil {
		log.Fatal(
			"failed to read config file",
			zap.Error(err),
		)
	}
	
	app := LoadAppConfig(vpr)
	db := LoadDatabaseConfig(vpr, log)
	redis := LoadRedisConfig(ctx, vpr, log)
	
	return &model.ConfigModel{
		App: app,
		Database: db,
		Redis: redis,
	}
}