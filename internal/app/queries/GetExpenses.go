package queries

import (
	"time"

	"github.com/pvtalbot/money/domain/models"
)

type GetExpenses struct {
	user      *models.User
	startDate time.Time
	endDate   time.Time
}

type GetExpensesQueryHandler struct {
	r models.ExpenseRepository
}

func NewGetExpensesHandler(r models.ExpenseRepository) GetExpensesQueryHandler {
	return GetExpensesQueryHandler{
		r: r,
	}
}

func NewGetExpensesQuery(user *models.User, startDate, endDate time.Time) GetExpenses {
	roundedStartDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	roundedEndDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day()+1, 0, 0, 0, 0, endDate.Location())

	return GetExpenses{
		user:      user,
		startDate: roundedStartDate,
		endDate:   roundedEndDate,
	}
}

func (h GetExpensesQueryHandler) Handle(q GetExpenses) ([]*models.Expense, error) {
	return h.r.GetAllExpensesFromUserBetweenDates(q.user, q.startDate, q.endDate)
}
