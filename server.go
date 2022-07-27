package main

import (
	"back_go/graph"
	"back_go/graph/generated"
	"back_go/internal/domain/managers"
	"back_go/internal/domain/model"
	"back_go/internal/infra/repositories"
	"back_go/internal/middlewares"

	database "back_go/internal/pkg/db/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func graphqlHandler(u model.UserServiceInterface, e model.ExpenseServiceInterface) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserService:    u,
		ExpenseService: e,
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
	e := managers.NewExpenseManager(repositories.NewExpenseMariaRepository(dbContainer.Db))

	r := gin.Default()

	r.Use(middlewares.GinContextToContextMiddleware())
	r.Use(middlewares.AuthMiddleware(u))

	r.POST("/query", graphqlHandler(u, e))
	r.GET("/", playgroundHandler())
	r.Run()
}
