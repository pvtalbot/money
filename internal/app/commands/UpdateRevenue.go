package commands

import (
	"log"
	"time"

	"github.com/pvtalbot/money/domain/models"
)

type UpdateRevenue struct {
	amount int
	date   time.Time
	id     string
}

type UpdateRevenueHandler struct {
	r models.RevenueRepository
}

func NewUpdateRevenueHandler(r models.RevenueRepository) UpdateRevenueHandler {
	return UpdateRevenueHandler{
		r: r,
	}
}

func NewUpdateRevenueCommand(revenue models.Revenue, amount *int, date *time.Time) UpdateRevenue {
	cmd := UpdateRevenue{id: revenue.ID}

	if amount == nil {
		cmd.amount = revenue.Amount
	} else {
		cmd.amount = *amount
	}

	if date == nil {
		cmd.date = revenue.GetDate()
	} else {
		cmd.date = *date
	}

	return cmd
}

func (h UpdateRevenueHandler) Handle(cmd UpdateRevenue) (*models.Revenue, error) {
	r := &models.Revenue{Amount: cmd.amount, ID: cmd.id}
	r.SetDate(cmd.date)

	log.Println("Update Revenue Command Handler", r.Amount, r.ID, r.GetDate())

	return h.r.Update(r)
}
