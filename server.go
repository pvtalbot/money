package main

import (
	"back_go/graph"
	"back_go/graph/generated"
	"back_go/internal/auth"
	"back_go/internal/domain/managers"
	"back_go/internal/domain/model"
	"back_go/internal/infra/repositories"

	database "back_go/internal/pkg/db/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func graphqlHandler(u model.UserServiceInterface) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserService: u,
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

	u := managers.NewUserManager(repositories.NewUserMariaRepository(dbContainer.Db))

	r := gin.Default()
	r.Use(auth.AuthMiddleware(u))
	r.POST("/query", graphqlHandler(u))
	r.GET("/", playgroundHandler())
	r.Run()
}
