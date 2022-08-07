package domain

import (
	"github.com/pvtalbot/money/app/domain/commands"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateExpense commands.CreateExpenseHandler
}

type Queries struct {
}
