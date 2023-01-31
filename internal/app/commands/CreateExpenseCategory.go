package commands

import "github.com/pvtalbot/money/domain/models"

type CreateExpenseCategory struct {
	Name   string
	UserId string
}

type CreateExpenseCategoryHandler struct {
	r models.ExpenseCategoryRepository
}

func NewCreateExpenseCategoryHandler(r models.ExpenseCategoryRepository) CreateExpenseCategoryHandler {
	return CreateExpenseCategoryHandler{
		r: r,
	}
}

func (h CreateExpenseCategoryHandler) Handle(cmd CreateExpenseCategory) (*models.ExpenseCategory, error) {
	expenseCategory := &models.ExpenseCategory{Name: cmd.Name}

	return h.r.Create(expenseCategory, cmd.UserId)
}
