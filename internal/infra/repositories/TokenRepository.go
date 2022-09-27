package repositories

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/pvtalbot/money/domain/models"
	"github.com/pvtalbot/money/pkg/jwt"
)

type TokenRedisRepository struct {
	redis *redis.Client
	ctx   context.Context
}

func NewTokenRedisRepository(redis *redis.Client) TokenRedisRepository {
	return TokenRedisRepository{
		redis: redis,
		ctx:   context.Background(),
	}
}

func (r TokenRedisRepository) Create(userId, userName string) (*models.Token, error) {
	authToken, err := jwt.GenerateToken(userName, userId)
	if err != nil {
		return nil, err
	}

	refreshToken := randSeq(100)

	key := "user:" + userId
	r.redis.HSet(r.ctx, key, "lastRefreshToken", refreshToken)
	r.redis.Expire(r.ctx, key, 7*24*time.Hour)

	return &models.Token{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r TokenRedisRepository) Find(userId string) (*models.Token, error) {
	key := "user:" + userId
	result := r.redis.HGet(r.ctx, key, "lastRefreshToken")

	if err := result.Err(); err != nil {
		return nil, err
	}

	return &models.Token{
		RefreshToken: result.Val(),
	}, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
