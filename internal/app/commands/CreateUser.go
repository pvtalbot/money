package commands

import (
	"github.com/pvtalbot/money/domain/managers"
	"github.com/pvtalbot/money/domain/models"
)

type CreateUser struct {
	user *models.User
}

type CreateUserHandler struct {
	r  models.UserRepository
	ec models.ExpenseCategoryRepository
}

func NewCreateUserHandler(r models.UserRepository, ec models.ExpenseCategoryRepository) CreateUserHandler {
	return CreateUserHandler{
		r:  r,
		ec: ec,
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
	user, err := h.r.Create(cmd.user)
	if err != nil {
		return nil, err
	}

	categories := managers.GetDefaultCategories(user)
	for _, c := range categories {
		_, err = h.ec.Create(&c)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
