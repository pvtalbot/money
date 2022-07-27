package managers

import "back_go/internal/domain/model"

type ExpenseManager struct {
	r model.ExpenseRepository
}

func NewExpenseManager(r model.ExpenseRepository) ExpenseManager {
	return ExpenseManager{
		r: r,
	}
}

func (m ExpenseManager) Create(amount int, user *model.User) *model.Expense {
	exp := &model.Expense{
		Amount: amount,
	}

	return m.r.Create(exp, user)
}
