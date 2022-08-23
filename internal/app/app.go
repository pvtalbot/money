package app

import (
	"database/sql"
	"os"

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

	// Revenues
	CreateRevenue commands.CreateRevenueHandler
	DeleteRevenue commands.DeleteRevenueHandler
	UpdateRevenue commands.UpdateRevenueHandler

	// Users
	CreateUser commands.CreateUserHandler
	Login      commands.LoginHandler
}

type Queries struct {
	// Errors
	GetAllErrors queries.GetAllErrorsHandler

	// Expenses
	FindExpense queries.FindExpenseQueryHandler
	GetExpenses queries.GetExpensesQueryHandler
	SumExpenses queries.SumExpensesQueryHandler

	// Expenses Categories
	GetExpensesCategories queries.GetExpensesCategoriesQueryHandler
	FindExpenseCategory   queries.FindExpenseCategoryQueryHandler

	// Revenues
	FindRevenue queries.FindRevenueQueryHandler
	GetRevenues queries.GetRevenuesQueryHandler

	// Users
	FindUser queries.FindUserQueryHandler
}

func NewApplication() (Application, func()) {
	dbContainer := database.NewDbContainer()
	if os.Getenv("MODE") == "dev" {
		dbContainer.Migrate()
	}

	return newApplication(dbContainer.Db),
		func() {
			_ = dbContainer.CloseDB()
		}
}

func newApplication(db *sql.DB) Application {
	expenseRepository := repositories.NewExpenseMariaRepository(db)
	expenseCategoryRepository := repositories.NewExpenseCategoryMariaRepository(db)
	revenueRepository := repositories.NewRevenueMariaRepository(db)
	userRepository := repositories.NewUserMariaRepository(db)

	return Application{
		Commands: Commands{
			// Expenses
			CreateExpense: commands.NewCreateExpenseHandler(expenseRepository),
			DeleteExpense: commands.NewDeleteExpenseHandler(expenseRepository),
			UpdateExpense: commands.NewUpdateExpenseHandler(expenseRepository),

			// Revenues
			CreateRevenue: commands.NewCreateRevenueHandler(revenueRepository),
			DeleteRevenue: commands.NewDeleteRevenueHandler(revenueRepository),
			UpdateRevenue: commands.NewUpdateRevenueHandler(revenueRepository),

			// Users
			CreateUser: commands.NewCreateUserHandler(userRepository, expenseCategoryRepository),
			Login:      commands.NewLoginHandler(userRepository),
		},
		Queries: Queries{
			// Errors
			GetAllErrors: queries.GetAllErrorsHandler{},

			// Expenses
			FindExpense: queries.NewFindExpenseHandler(expenseRepository),
			GetExpenses: queries.NewGetExpensesHandler(expenseRepository),
			SumExpenses: queries.NewSumExpensesQueryHandler(expenseRepository),

			// Expenses Categories
			FindExpenseCategory:   queries.NewFindExpenseCategoryHandler(expenseCategoryRepository),
			GetExpensesCategories: queries.NewGetExpensesCategoriesHandler(expenseCategoryRepository),

			// Revenues
			FindRevenue: queries.NewFindRevenueHandler(revenueRepository),
			GetRevenues: queries.NewGetRevenuesHandler(revenueRepository),

			// Users
			FindUser: queries.NewFindUserHandler(userRepository),
		},
	}
}
