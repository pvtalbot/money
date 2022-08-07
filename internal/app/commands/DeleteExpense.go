package commands

import (
	"github.com/pvtalbot/money/domain/model"
)

type DeleteExpense struct {
	Id string
}

type DeleteExpenseHandler struct {
	r model.ExpenseRepository
}

func NewDeleteExpenseHandler(r model.ExpenseRepository) DeleteExpenseHandler {
	return DeleteExpenseHandler{
		r: r,
	}
}

func (h DeleteExpenseHandler) Handle(cmd DeleteExpense) error {
	return h.r.Delete(cmd.Id)
}
