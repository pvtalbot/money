package domain

import (
	"github.com/pvtalbot/money/app/domain/commands"
	"github.com/pvtalbot/money/app/domain/queries"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateExpense commands.CreateExpenseHandler
}

type Queries struct {
	FindExpense queries.FindExpenseQueryHandler
	GetExpenses queries.GetExpensesQueryHandler
}
