package service

import (
	"database/sql"

	"github.com/pvtalbot/money/app"
	"github.com/pvtalbot/money/app/commands"
	"github.com/pvtalbot/money/app/queries"
	"github.com/pvtalbot/money/infra/repositories"

	database "github.com/pvtalbot/money/pkg/db/mysql"
)

func NewApplication() (app.Application, func()) {
	dbContainer := database.NewDbContainer()
	dbContainer.Migrate()

	return newApplication(dbContainer.Db),
		func() {
			_ = dbContainer.CloseDB()
		}
}

func newApplication(db *sql.DB) app.Application {
	expenseRepository := repositories.NewExpenseMariaRepository(db)
	userRepository := repositories.NewUserMariaRepository(db)

	return app.Application{
		Commands: app.Commands{
			// Expenses
			CreateExpense: commands.NewCreateExpenseHandler(expenseRepository),
			DeleteExpense: commands.NewDeleteExpenseHandler(expenseRepository),
			UpdateExpense: commands.NewUpdateExpenseHandler(expenseRepository),

			// Users
			CreateUser: commands.NewCreateUserHandler(userRepository),
			Login:      commands.NewLoginHandler(userRepository),
		},
		Queries: app.Queries{
			// Expenses
			FindExpense: queries.NewFindExpenseHandler(expenseRepository),
			GetExpenses: queries.NewGetExpensesHandler(expenseRepository),
			SumExpenses: queries.NewSumExpensesQueryHandler(expenseRepository),

			// Users
			FindUser: queries.NewFindUserHandler(userRepository),
		},
	}
}
