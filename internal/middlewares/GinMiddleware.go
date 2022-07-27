package middlewares

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

const CONTEXT_KEY_NAME string = "GinContextKey"

type ginContextKey struct {
	name string
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	key := ginContextKey{name: CONTEXT_KEY_NAME}

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), key, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	key := ginContextKey{name: CONTEXT_KEY_NAME}

	ginContext := ctx.Value(key)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)

	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}

	return gc, nil
}
