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
}

func NewTokenRedisRepository(redis *redis.Client) TokenRedisRepository {
	return TokenRedisRepository{
		redis: redis,
	}
}

func (r TokenRedisRepository) Create(userId, userName string) (*models.Token, error) {
	authToken, err := jwt.GenerateToken(userName, userId)
	if err != nil {
		return nil, err
	}

	refreshToken := randSeq(100)

	ctx := context.Background()
	key := "user:" + userId
	r.redis.HSet(ctx, key, "lastRefreshToken", refreshToken)
	r.redis.Expire(ctx, key, 7*24*time.Hour)

	return &models.Token{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
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
