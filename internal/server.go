package main

import (
	"github.com/pvtalbot/money/app"
	"github.com/pvtalbot/money/graph"
	"github.com/pvtalbot/money/graph/generated"
	"github.com/pvtalbot/money/middlewares"
	"github.com/pvtalbot/money/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

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

func main() {
	app, cleanup := service.NewApplication()
	defer cleanup()

	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.GinContextToContextMiddleware())
	r.Use(middlewares.AuthMiddleware())

	r.POST("/query", graphqlHandler(app))
	r.GET("/", playgroundHandler())
	r.Run()
}
