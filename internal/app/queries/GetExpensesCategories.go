package queries

import (
	"github.com/pvtalbot/money/domain/models"
)

type GetExpensesCategories struct {
	User *models.User
}

type GetExpensesCategoriesQueryHandler struct {
	r models.ExpenseCategoryRepository
}

func NewGetExpensesCategoriesHandler(r models.ExpenseCategoryRepository) GetExpensesCategoriesQueryHandler {
	return GetExpensesCategoriesQueryHandler{
		r: r,
	}
}

func (h GetExpensesCategoriesQueryHandler) Handle(q GetExpensesCategories) ([]*models.ExpenseCategory, error) {
	return h.r.FindAll(q.User)
}
