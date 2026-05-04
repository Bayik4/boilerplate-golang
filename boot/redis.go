package boot

import (
	"context"
	"fmt"

	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func RedisInit(ctx context.Context, cfg model.RedisConfig, log *zap.Logger) *redis.Client {
	addr := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DbIndex,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed connect to redis", zap.String("addr", addr))
	}

	log.Info("Redis client connected", zap.String("addr", addr))

	return rdb
}
