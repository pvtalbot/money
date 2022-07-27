package middlewares

import (
	"back_go/internal/domain/managers"
	"back_go/internal/domain/model"
	"back_go/pkg/jwt"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(um managers.UserManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		if header == "" {
			c.Next()
			return
		}

		tokenStr := header
		username, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := um.FindByName(username)
		if err != nil {
			c.Next()
		}

		c.Set("user", user)
		c.Next()
	}
}

func ExtractUserFromContext(ctx context.Context) (*model.User, error) {
	ginContext, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	u, ok := ginContext.Get("user")
	if !ok {
		return nil, errors.New("not found")
	}

	user, ok := u.(*model.User)
	if !ok {
		return nil, errors.New("user has wrong type")
	}

	return user, nil
}
