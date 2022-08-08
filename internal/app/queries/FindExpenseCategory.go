package queries

import "github.com/pvtalbot/money/domain/models"

type FindExpenseCategory struct {
	Id string
}

type FindExpenseCategoryQueryHandler struct {
	r models.ExpenseCategoryRepository
}

func NewFindExpenseCategoryHandler(r models.ExpenseCategoryRepository) FindExpenseCategoryQueryHandler {
	return FindExpenseCategoryQueryHandler{
		r: r,
	}
}

func (h FindExpenseCategoryQueryHandler) Handle(q FindExpenseCategory) (*models.ExpenseCategory, error) {
	return h.r.Find(q.Id)
}
