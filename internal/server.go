package main

import (
	"github.com/pvtalbot/money/app/domain/managers"
	"github.com/pvtalbot/money/app/domain/model"
	"github.com/pvtalbot/money/app/infra/repositories"
	"github.com/pvtalbot/money/app/middlewares"
	"github.com/pvtalbot/money/graph"
	"github.com/pvtalbot/money/graph/generated"

	database "github.com/pvtalbot/money/app/pkg/db/mysql"

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

	expenseRepository := repositories.NewExpenseMariaRepository(dbContainer.Db)
	userRepository := repositories.NewUserMariaRepository(dbContainer.Db)

	u := managers.NewUserManager(userRepository, expenseRepository)
	e := managers.NewExpenseManager(expenseRepository)

	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.GinContextToContextMiddleware())
	r.Use(middlewares.AuthMiddleware(u))

	r.POST("/query", graphqlHandler(u, e))
	r.GET("/", playgroundHandler())
	r.Run()
}
