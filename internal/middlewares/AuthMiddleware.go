package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/pvtalbot/money/domain/models"
	"github.com/pvtalbot/money/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		if header == "" {
			c.Next()
			return
		}

		tokenStr := header
		user, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func ExtractUserFromContext(ctx context.Context) (*models.User, error) {
	ginContext, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	u, ok := ginContext.Get("user")
	if !ok {
		return nil, errors.New("not found")
	}

	user, ok := u.(*models.User)
	if !ok {
		return nil, errors.New("user has wrong type")
	}

	return user, nil
}
