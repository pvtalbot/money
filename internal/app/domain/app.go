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
	//Expenses
	CreateExpense commands.CreateExpenseHandler
	DeleteExpense commands.DeleteExpenseHandler
	UpdateExpense commands.UpdateExpenseHandler

	//Users
	CreateUser commands.CreateUserHandler
}

type Queries struct {
	//Expenses
	FindExpense queries.FindExpenseQueryHandler
	GetExpenses queries.GetExpensesQueryHandler
	SumExpenses queries.SumExpensesQueryHandler
}
