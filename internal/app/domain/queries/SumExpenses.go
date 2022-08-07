package queries

import (
	"time"

	"github.com/pvtalbot/money/app/domain/managers"
	"github.com/pvtalbot/money/app/domain/model"
	"github.com/pvtalbot/money/pkg/utils"
)

type SumExpenses struct {
	startDate time.Time
	endDate   time.Time
	user      *model.User
	groupBy   utils.Duration
}

type SumExpensesQueryHandler struct {
	r model.ExpenseRepository
}

func NewSumExpensesQueryHandler(r model.ExpenseRepository) SumExpensesQueryHandler {
	return SumExpensesQueryHandler{
		r: r,
	}
}

func NewSumExpensesQuery(startDate, endDate time.Time, user *model.User, groupBy utils.Duration) SumExpenses {
	var roundedStartDate, roundedEndDate time.Time
	switch groupBy {
	case utils.MONTH:
		roundedStartDate = time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, startDate.Location())
		roundedEndDate = time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, endDate.Location())
	}

	return SumExpenses{
		startDate: roundedStartDate,
		endDate:   roundedEndDate,
		user:      user,
		groupBy:   groupBy,
	}
}

func (h SumExpensesQueryHandler) Handle(q SumExpenses) ([]*model.ExpenseSum, error) {
	var result []*model.ExpenseSum
	var err error
	switch q.groupBy {
	case utils.MONTH:
		if q.startDate.Year() == q.endDate.Year() && q.startDate.Month() == q.endDate.Month() {
			return result, nil
		}

		result, err = h.r.SumAllExpensesFromUserBetweenDatesByMonth(q.user, q.startDate, q.endDate)
		if err != nil {
			return nil, err
		}

		result = managers.PopulateResult(q.startDate, q.endDate, result)
	}

	return result, nil
}
