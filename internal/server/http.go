package server

import (
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/pvtalbot/money/app"
	"github.com/pvtalbot/money/graph"
	"github.com/pvtalbot/money/graph/generated"
	"github.com/pvtalbot/money/middlewares"
)

func RunHttpServer(app app.Application) {
	r := gin.Default()

	setMiddlewares(r)
	setHandlers(r, app)
	r.SetTrustedProxies(nil) // Traefik and Docker are here :)

	r.Run()
}

func setMiddlewares(r *gin.Engine) {
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.GinContextToContextMiddleware())
	r.Use(middlewares.AuthMiddleware())
}

func setHandlers(r *gin.Engine, app app.Application) {
	r.POST("/query", graphqlHandler(app))
	if os.Getenv("GIN_MODE") != "release" {
		r.GET("/", playgroundHandler())
	}
}

func graphqlHandler(app app.Application) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Application: app,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
