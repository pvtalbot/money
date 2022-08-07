package commands

import (
	"time"

	"github.com/pvtalbot/money/app/domain/model"
)

type CreateExpense struct {
	Amount int
	Date   time.Time
	User   *model.User
}

type CreateExpenseHandler struct {
	r model.ExpenseRepository
}

func NewCreateExpenseHandler(r model.ExpenseRepository) CreateExpenseHandler {
	return CreateExpenseHandler{
		r: r,
	}
}

func (h CreateExpenseHandler) Handle(cmd CreateExpense) (*model.Expense, error) {
	exp := &model.Expense{Amount: cmd.Amount}
	exp.SetDate(cmd.Date)

	return h.r.Create(exp, cmd.User)
}
