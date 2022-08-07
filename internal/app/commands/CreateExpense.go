package commands

import (
	"time"

	"github.com/pvtalbot/money/domain/models"
)

type CreateExpense struct {
	Amount int
	Date   time.Time
	User   *models.User
}

type CreateExpenseHandler struct {
	r models.ExpenseRepository
}

func NewCreateExpenseHandler(r models.ExpenseRepository) CreateExpenseHandler {
	return CreateExpenseHandler{
		r: r,
	}
}

func (h CreateExpenseHandler) Handle(cmd CreateExpense) (*models.Expense, error) {
	exp := &models.Expense{Amount: cmd.Amount}
	exp.SetDate(cmd.Date)

	return h.r.Create(exp, cmd.User)
}
