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

	return domain.Application{
		Commands: domain.Commands{
			CreateExpense: commands.NewCreateExpenseHandler(expenseRepository),
		},
		Queries: domain.Queries{
			FindExpense: queries.NewFindExpenseHandler(expenseRepository),
			GetExpenses: queries.NewGetExpensesHandler(expenseRepository),
		},
	}
}
