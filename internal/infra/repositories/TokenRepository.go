package repositories

import (
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

	return &models.Token{
		AuthToken:    authToken,
		RefreshToken: "1",
	}, nil
}
