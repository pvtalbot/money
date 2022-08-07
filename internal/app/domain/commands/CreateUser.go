package commands

import (
	"github.com/pvtalbot/money/app/domain/model"
)

type CreateUser struct {
	user *model.User
}

type CreateUserHandler struct {
	r model.UserRepository
}

func NewCreateUserHandler(r model.UserRepository) CreateUserHandler {
	return CreateUserHandler{
		r: r,
	}
}

func NewCreateUser(name, password, firstName, lastName string) CreateUser {
	u := &model.User{
		Name:      name,
		FirstName: firstName,
		LastName:  lastName,
	}
	u.SetPassword(password)

	return CreateUser{user: u}
}

func (h CreateUserHandler) Handle(cmd CreateUser) (*model.User, error) {
	return h.r.Create(cmd.user)
}
