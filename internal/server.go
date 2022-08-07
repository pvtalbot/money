package main

import (
	"github.com/pvtalbot/money/app/domain"
	"github.com/pvtalbot/money/app/middlewares"
	"github.com/pvtalbot/money/app/service"
	"github.com/pvtalbot/money/graph"
	"github.com/pvtalbot/money/graph/generated"

	database "github.com/pvtalbot/money/app/pkg/db/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func graphqlHandler(app domain.Application) gin.HandlerFunc {
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
	dbContainer := database.NewDbContainer()
	defer dbContainer.CloseDB()
	dbContainer.Migrate()

	app := service.NewApplication(dbContainer.Db)

	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.GinContextToContextMiddleware())
	r.Use(middlewares.AuthMiddleware())

	r.POST("/query", graphqlHandler(app))
	r.GET("/", playgroundHandler())
	r.Run()
}
