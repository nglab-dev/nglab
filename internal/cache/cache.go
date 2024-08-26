package cache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/pkg/env"
	"github.com/nglab-dev/nglab/pkg/log"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Memory *bigcache.BigCache
	Redis  *redis.Client
}

func Init(c *config.Config) (*Cache, error) {

	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		return nil, err
	}

	log.Logger.Sugar().Info("Memory cache initialized")

	redisEnalbed, _ := env.GetBool("REDIS_ENABLED", false)
	var rdb *redis.Client
	if redisEnalbed {
		redisAddr := env.GetString("REDIS_ADDR", "localhost:6379")
		redisPassword := env.GetString("REDIS_PASSWORD", "")
		redisDB, _ := env.GetInt("REDIS_DB", 0)
		rdb = redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       redisDB,
		})
		log.Logger.Sugar().Info("Redis cache initialized")
	}

	return &Cache{Memory: cache, Redis: rdb}, nil
}
