package queries

import "github.com/pvtalbot/money/domain/models"

type FindExpense struct {
	Id string
}

type FindExpenseQueryHandler struct {
	r models.ExpenseRepository
}

func NewFindExpenseHandler(r models.ExpenseRepository) FindExpenseQueryHandler {
	return FindExpenseQueryHandler{
		r: r,
	}
}

func (h FindExpenseQueryHandler) Handle(q FindExpense) (*models.Expense, error) {
	return h.r.Find(q.Id)
}
