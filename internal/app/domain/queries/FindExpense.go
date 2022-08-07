package queries

import "github.com/pvtalbot/money/app/domain/model"

type FindExpense struct {
	Id string
}

type FindExpenseQueryHandler struct {
	r model.ExpenseRepository
}

func NewFindExpenseHandler(r model.ExpenseRepository) FindExpenseQueryHandler {
	return FindExpenseQueryHandler{
		r: r,
	}
}

func (h FindExpenseQueryHandler) Handle(q FindExpense) (*model.Expense, error) {
	return h.r.Find(q.Id)
}
