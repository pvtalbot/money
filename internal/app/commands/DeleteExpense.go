package commands

import (
	"github.com/pvtalbot/money/domain/models"
)

type DeleteExpense struct {
	Id string
}

type DeleteExpenseHandler struct {
	r models.ExpenseRepository
}

func NewDeleteExpenseHandler(r models.ExpenseRepository) DeleteExpenseHandler {
	return DeleteExpenseHandler{
		r: r,
	}
}

func (h DeleteExpenseHandler) Handle(cmd DeleteExpense) error {
	return h.r.Delete(cmd.Id)
}
