package commands

import (
	"time"

	"github.com/pvtalbot/money/domain/models"
)

type UpdateExpense struct {
	amount int
	date   time.Time
	id     string
}

type UpdateExpenseHandler struct {
	r models.ExpenseRepository
}

func NewUpdateExpenseHandler(r models.ExpenseRepository) UpdateExpenseHandler {
	return UpdateExpenseHandler{
		r: r,
	}
}

func NewUpdateExpenseCommand(expense models.Expense, amount *int, date *time.Time) UpdateExpense {
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

func (h UpdateExpenseHandler) Handle(cmd UpdateExpense) (*models.Expense, error) {
	exp := &models.Expense{Amount: cmd.amount, ID: cmd.id}
	exp.SetDate(cmd.date)

	return h.r.Update(exp)
}
