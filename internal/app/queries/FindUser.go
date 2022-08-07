package queries

import "github.com/pvtalbot/money/domain/model"

type FindUser struct {
	Id string
}

type FindUserQueryHandler struct {
	r model.UserRepository
}

func NewFindUserHandler(r model.UserRepository) FindUserQueryHandler {
	return FindUserQueryHandler{
		r: r,
	}
}

func (h FindUserQueryHandler) Handle(q FindUser) (*model.User, error) {
	return h.r.Find(q.Id)
}
