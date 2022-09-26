package redis

import (
	"os"

	"github.com/go-redis/redis/v9"
)

type RedisClientContainer struct {
	Redis *redis.Client
}

func NewRedisContainer() *RedisClientContainer {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	r := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	return &RedisClientContainer{
		Redis: r,
	}
}

func (r *RedisClientContainer) CloseRedisClient() error {
	return r.Redis.Close()
}
