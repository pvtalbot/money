package auth

import (
	"back_go/internal/pkg/users"
	"back_go/pkg/jwt"
	"net/http"

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
		username, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := users.GetUserByName(username)
		if err != nil {
			c.Next()
		}

		c.Set("user", &user)
		c.Next()
	}
}
