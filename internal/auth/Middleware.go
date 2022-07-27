package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		/*header := c.Request.Header.Get("Authorization")

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
		c.Next()*/
	}
}
