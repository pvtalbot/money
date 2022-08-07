package queries

import "github.com/pvtalbot/money/domain/models"

type FindUser struct {
	Id string
}

type FindUserQueryHandler struct {
	r models.UserRepository
}

func NewFindUserHandler(r models.UserRepository) FindUserQueryHandler {
	return FindUserQueryHandler{
		r: r,
	}
}

func (h FindUserQueryHandler) Handle(q FindUser) (*models.User, error) {
	return h.r.Find(q.Id)
}
