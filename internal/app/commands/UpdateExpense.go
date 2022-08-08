package commands

import (
	"time"

	"github.com/pvtalbot/money/domain/models"
)

type UpdateExpense struct {
	amount     int
	date       time.Time
	categoryId string
	id         string
}

type UpdateExpenseHandler struct {
	r models.ExpenseRepository
}

func NewUpdateExpenseHandler(r models.ExpenseRepository) UpdateExpenseHandler {
	return UpdateExpenseHandler{
		r: r,
	}
}

func NewUpdateExpenseCommand(expense models.Expense, amount *int, date *time.Time, categoryId *string) UpdateExpense {
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

	if categoryId == nil {
		cmd.categoryId = expense.Category.ID
	} else {
		cmd.categoryId = *categoryId
	}

	return cmd
}

func (h UpdateExpenseHandler) Handle(cmd UpdateExpense) (*models.Expense, error) {
	exp := &models.Expense{Amount: cmd.amount, ID: cmd.id, Category: models.ExpenseCategory{ID: cmd.categoryId}}
	exp.SetDate(cmd.date)

	return h.r.Update(exp)
}
