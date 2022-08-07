package commands

import (
	"github.com/pvtalbot/money/domain/models"
)

type CreateUser struct {
	user *models.User
}

type CreateUserHandler struct {
	r models.UserRepository
}

func NewCreateUserHandler(r models.UserRepository) CreateUserHandler {
	return CreateUserHandler{
		r: r,
	}
}

func NewCreateUser(name, password, firstName, lastName string) CreateUser {
	u := &models.User{
		Name:      name,
		FirstName: firstName,
		LastName:  lastName,
	}
	u.SetPassword(password)

	return CreateUser{user: u}
}

func (h CreateUserHandler) Handle(cmd CreateUser) (*models.User, error) {
	return h.r.Create(cmd.user)
}
