package app

import (
	"database/sql"

	"github.com/pvtalbot/money/app/commands"
	"github.com/pvtalbot/money/app/queries"
	"github.com/pvtalbot/money/infra/repositories"

	database "github.com/pvtalbot/money/pkg/db/mysql"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	// Expenses
	CreateExpense commands.CreateExpenseHandler
	DeleteExpense commands.DeleteExpenseHandler
	UpdateExpense commands.UpdateExpenseHandler

	// Users
	CreateUser commands.CreateUserHandler
	Login      commands.LoginHandler
}

type Queries struct {
	//Expenses
	FindExpense queries.FindExpenseQueryHandler
	GetExpenses queries.GetExpensesQueryHandler
	SumExpenses queries.SumExpensesQueryHandler

	// Users
	FindUser queries.FindUserQueryHandler
}

func NewApplication() (Application, func()) {
	dbContainer := database.NewDbContainer()
	dbContainer.Migrate()

	return newApplication(dbContainer.Db),
		func() {
			_ = dbContainer.CloseDB()
		}
}

func newApplication(db *sql.DB) Application {
	expenseRepository := repositories.NewExpenseMariaRepository(db)
	expenseCategoryRepository := repositories.NewExpenseCategoryMariaRepository(db)
	userRepository := repositories.NewUserMariaRepository(db)

	return Application{
		Commands: Commands{
			//Expenses
			CreateExpense: commands.NewCreateExpenseHandler(expenseRepository),
			DeleteExpense: commands.NewDeleteExpenseHandler(expenseRepository),
			UpdateExpense: commands.NewUpdateExpenseHandler(expenseRepository),

			//Users
			CreateUser: commands.NewCreateUserHandler(userRepository, expenseCategoryRepository),
			Login:      commands.NewLoginHandler(userRepository),
		},
		Queries: Queries{
			//Expenses
			FindExpense: queries.NewFindExpenseHandler(expenseRepository),
			GetExpenses: queries.NewGetExpensesHandler(expenseRepository),
			SumExpenses: queries.NewSumExpensesQueryHandler(expenseRepository),

			//Users
			FindUser: queries.NewFindUserHandler(userRepository),
		},
	}
}
