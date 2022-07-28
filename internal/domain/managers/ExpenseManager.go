package managers

import (
	"back_go/internal/domain/model"
	"time"
)

type ExpenseManager struct {
	r model.ExpenseRepository
}

func NewExpenseManager(r model.ExpenseRepository) ExpenseManager {
	return ExpenseManager{
		r: r,
	}
}

func (m ExpenseManager) Create(amount int, date time.Time, user *model.User) *model.Expense {
	roundedDate := time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, date.Location())
	exp := &model.Expense{
		Amount: amount,
		Date:   roundedDate,
	}

	return m.r.Create(exp, user)
}
