package commands

import "github.com/pvtalbot/money/domain/models"

type DeleteRevenue struct {
	Id string
}

type DeleteRevenueHandler struct {
	r models.RevenueRepository
}

func NewDeleteRevenueHandler(r models.RevenueRepository) DeleteRevenueHandler {
	return DeleteRevenueHandler{
		r: r,
	}
}

func (h DeleteRevenueHandler) Handle(cmd DeleteRevenue) error {
	return h.r.Delete(cmd.Id)
}
