package commands

import (
	"time"

	"github.com/pvtalbot/money/app/domain/model"
)

type UpdateExpense struct {
	amount int
	date   time.Time
	id     string
}

type UpdateExpenseHandler struct {
	r model.ExpenseRepository
}

func NewUpdateExpenseHandler(r model.ExpenseRepository) UpdateExpenseHandler {
	return UpdateExpenseHandler{
		r: r,
	}
}

func NewUpdateExpenseCommand(expense model.Expense, amount *int, date *time.Time) UpdateExpense {
	cmd := UpdateExpense{id: expense.ID}

	if amount == nil {
		cmd.amount = expense.Amount
	} else {
		cmd.amount = *amount
	}

	if date == nil {
		cmd.date = expense.GetDate()
	} else {
		cmd.date = *date
	}

	return cmd
}

func (h UpdateExpenseHandler) Handle(cmd UpdateExpense) (*model.Expense, error) {
	exp := &model.Expense{Amount: cmd.amount, ID: cmd.id}
	exp.SetDate(cmd.date)

	return h.r.Update(exp)
}
