package redis

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v9"
)

type RedisClientContainer struct {
	Redis *redis.Client
	Ctx   context.Context
}

func NewRedisContainer() *RedisClientContainer {
	ctx := context.Background()

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	r := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	if err := r.Ping(ctx).Err(); err != nil {
		log.Panic(err)
	}

	return &RedisClientContainer{
		Redis: r,
	}
}

func (r *RedisClientContainer) CloseRedisClient() error {
	return r.Redis.Close()
}
