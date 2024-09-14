package keystorefx

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RedisParams struct {
	fx.In

	Config KeyStoreConfig
	Logger *zap.Logger
}

func NewRedisClient(p RedisParams) *redis.Client {
	cfg := p.Config
	opts := &redis.Options{
		Network: "tcp",
		Addr:    fmt.Sprint(cfg.Address, ":", cfg.Port),
		Password: cfg.Password,
	}
	redisConnection := redis.NewClient(opts)

	res := redisConnection.Ping(context.TODO())
	if res.Err() != nil {
		log.Print("could not connect to redis, %w", res.Err())
	}

	return redisConnection
}
