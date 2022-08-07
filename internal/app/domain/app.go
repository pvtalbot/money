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
	DeleteExpense commands.DeleteExpenseHandler
	UpdateExpense commands.UpdateExpenseHandler
}

type Queries struct {
	FindExpense queries.FindExpenseQueryHandler
	GetExpenses queries.GetExpensesQueryHandler
	SumExpenses queries.SumExpensesQueryHandler
}
