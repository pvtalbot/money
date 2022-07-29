package managers

import (
	"back_go/internal/domain/model"
	"errors"
	"strconv"
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

func (m ExpenseManager) GetAllExpensesFromUserBetweenDates(user *model.User, startDate, endDate time.Time) []*model.Expense {
	roundedStartDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	roundedEndDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day()+1, 0, 0, 0, 0, endDate.Location())

	return m.r.GetAllExpensesFromUserBetweenDates(user, roundedStartDate, roundedEndDate)
}

func (m ExpenseManager) Delete(id, userID string) (*model.Expense, error) {
	intId, _ := strconv.ParseInt(id, 10, 64)

	expense, err := m.r.Find(intId)

	if err != nil {
		return nil, err
	}

	if userID != expense.User.ID {
		return nil, errors.New("user cannot delete expense")
	}

	return expense, m.r.Delete(intId)
}
