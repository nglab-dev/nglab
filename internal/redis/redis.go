package redis

import (
	"context"
	"log/slog"
	"time"

	"github.com/nglab-dev/nglab/internal/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func New(cfg config.Config) Redis {
	addr := cfg.Redis.Addr()

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       cfg.Redis.DB,
		Password: cfg.Redis.Password,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		slog.Error("Redis connection failed", "error", err)
	}

	slog.Info("Redis connection established")

	return Redis{client}
}
