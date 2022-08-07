package service

import (
	"database/sql"

	"github.com/pvtalbot/money/app/domain"
	"github.com/pvtalbot/money/app/domain/commands"
	"github.com/pvtalbot/money/app/domain/queries"
	"github.com/pvtalbot/money/app/infra/repositories"
)

func NewApplication(db *sql.DB) domain.Application {
	return newApplication(db)
}

func newApplication(db *sql.DB) domain.Application {
	expenseRepository := repositories.NewExpenseMariaRepository(db)
	userRepository := repositories.NewUserMariaRepository(db)

	return domain.Application{
		Commands: domain.Commands{
			// Expenses
			CreateExpense: commands.NewCreateExpenseHandler(expenseRepository),
			DeleteExpense: commands.NewDeleteExpenseHandler(expenseRepository),
			UpdateExpense: commands.NewUpdateExpenseHandler(expenseRepository),

			// Users
			CreateUser: commands.NewCreateUserHandler(userRepository),
			Login:      commands.NewLoginHandler(userRepository),
		},
		Queries: domain.Queries{
			// Expenses
			FindExpense: queries.NewFindExpenseHandler(expenseRepository),
			GetExpenses: queries.NewGetExpensesHandler(expenseRepository),
			SumExpenses: queries.NewSumExpensesQueryHandler(expenseRepository),

			// Users
			FindUser: queries.NewFindUserHandler(userRepository),
		},
	}
}
