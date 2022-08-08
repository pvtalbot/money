package commands

import (
	"time"

	"github.com/pvtalbot/money/domain/models"
)

type CreateRevenue struct {
	Amount int
	Date   time.Time
	UserId string
}

type CreateRevenueHandler struct {
	r models.RevenueRepository
}

func NewCreateRevenueHandler(r models.RevenueRepository) CreateRevenueHandler {
	return CreateRevenueHandler{
		r: r,
	}
}

func (h CreateRevenueHandler) Handle(cmd CreateRevenue) (*models.Revenue, error) {
	r := &models.Revenue{Amount: cmd.Amount}
	r.SetDate(cmd.Date)

	return h.r.Create(r, cmd.UserId)
}
