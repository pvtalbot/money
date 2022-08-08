package queries

import "github.com/pvtalbot/money/domain/models"

type FindRevenue struct {
	Id string
}

type FindRevenueQueryHandler struct {
	r models.RevenueRepository
}

func NewFindRevenueHandler(r models.RevenueRepository) FindRevenueQueryHandler {
	return FindRevenueQueryHandler{
		r: r,
	}
}

func (h FindRevenueQueryHandler) Handle(q FindRevenue) (*models.Revenue, error) {
	return h.r.Find(q.Id)
}
