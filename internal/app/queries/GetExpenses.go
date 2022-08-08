package queries

import (
	"time"

	"github.com/pvtalbot/money/domain/models"
)

type GetExpenses struct {
	userId    string
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

func NewGetExpensesQuery(userId string, startDate, endDate time.Time) GetExpenses {
	roundedStartDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	roundedEndDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day()+1, 0, 0, 0, 0, endDate.Location())

	return GetExpenses{
		userId:    userId,
		startDate: roundedStartDate,
		endDate:   roundedEndDate,
	}
}

func (h GetExpensesQueryHandler) Handle(q GetExpenses) ([]*models.Expense, error) {
	return h.r.GetAllExpensesFromUserBetweenDates(q.userId, q.startDate, q.endDate)
}
