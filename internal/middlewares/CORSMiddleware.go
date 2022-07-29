package middlewares

import (
	"back_go/pkg/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	frontHost, err := config.GetFrontHost()
	if err != nil {
		frontHost = "*"
	}

	config := cors.Config{
		AllowOrigins:     []string{frontHost},
		AllowMethods:     []string{"GET", "POST", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{},
		MaxAge:           12 * time.Hour,
	}

	return cors.New(config)
}
