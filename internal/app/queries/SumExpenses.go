package queries

import (
	"time"

	"github.com/pvtalbot/money/domain/managers"
	"github.com/pvtalbot/money/domain/models"
	"github.com/pvtalbot/money/pkg/utils"
)

type SumExpenses struct {
	startDate time.Time
	endDate   time.Time
	userId    string
	groupBy   utils.Duration
}

type SumExpensesQueryHandler struct {
	r models.ExpenseRepository
}

func NewSumExpensesQueryHandler(r models.ExpenseRepository) SumExpensesQueryHandler {
	return SumExpensesQueryHandler{
		r: r,
	}
}

func NewSumExpensesQuery(startDate, endDate time.Time, userId string, groupBy utils.Duration) SumExpenses {
	var roundedStartDate, roundedEndDate time.Time
	switch groupBy {
	case utils.MONTH:
		roundedStartDate = time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, startDate.Location())
		roundedEndDate = time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, endDate.Location())
	}

	return SumExpenses{
		startDate: roundedStartDate,
		endDate:   roundedEndDate,
		userId:    userId,
		groupBy:   groupBy,
	}
}

func (h SumExpensesQueryHandler) Handle(q SumExpenses) ([]*models.ExpenseSum, error) {
	var result []*models.ExpenseSum
	var err error
	switch q.groupBy {
	case utils.MONTH:
		if q.startDate.Year() == q.endDate.Year() && q.startDate.Month() == q.endDate.Month() {
			return result, nil
		}

		result, err = h.r.SumAllExpensesFromUserBetweenDatesByMonth(q.userId, q.startDate, q.endDate)
		if err != nil {
			return nil, err
		}

		result = managers.PopulateExpensesSum(q.startDate, q.endDate, result)
	}

	return result, nil
}
