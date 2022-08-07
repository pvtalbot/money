package app

import (
	"github.com/pvtalbot/money/app/commands"
	"github.com/pvtalbot/money/app/queries"
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
