package queries

import (
	"time"

	"github.com/pvtalbot/money/domain/models"
)

type GetRevenues struct {
	user      *models.User
	startDate time.Time
	endDate   time.Time
}

type GetRevenuesQueryHandler struct {
	r models.RevenueRepository
}

func NewGetRevenuesHandler(r models.RevenueRepository) GetRevenuesQueryHandler {
	return GetRevenuesQueryHandler{
		r: r,
	}
}

func NewGetRevenuesQuery(user *models.User, startDate, endDate time.Time) GetRevenues {
	roundedStartDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	roundedEndDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day()+1, 0, 0, 0, 0, endDate.Location())

	return GetRevenues{
		user:      user,
		startDate: roundedStartDate,
		endDate:   roundedEndDate,
	}
}

func (h GetRevenuesQueryHandler) Handle(q GetRevenues) ([]*models.Revenue, error) {
	return h.r.GetAllRevenuesOfUserBetweenDates(q.user, q.startDate, q.endDate)
}
