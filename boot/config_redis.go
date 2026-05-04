package boot

import (
	"context"

	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func LoadRedisConfig(ctx context.Context, vpr *viper.Viper, log *zap.Logger) model.RedisConfig {
	cfg := model.RedisConfig{
		Host: vpr.GetString("redis.host"),
		Port: vpr.GetString("redis.port"),
		DbIndex: vpr.GetInt("redis.dbindex"),
	}
	
	rdb := RedisInit(ctx, cfg, log)
	cfg.Conn = rdb
	
	return cfg
}