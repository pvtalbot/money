package managers

import (
	"back_go/internal/domain/model"
	"back_go/pkg/jwt"
	"back_go/pkg/utils"
)

type UserManager struct {
	r                 model.UserRepository
	expenseRepository model.ExpenseRepository
}

func NewUserManager(r model.UserRepository, expenseRepository model.ExpenseRepository) UserManager {
	return UserManager{
		r:                 r,
		expenseRepository: expenseRepository,
	}
}

func (m UserManager) Create(username, clearPassword, firstName, lastName string) *model.User {
	u := &model.User{
		Name:      username,
		FirstName: firstName,
		LastName:  lastName,
	}
	u.SetPassword(clearPassword)

	return m.r.Create(u)
}

func (m UserManager) FindAll() []*model.User {
	return m.r.FindAll()
}

func (m UserManager) FindByName(username string) (*model.User, error) {
	return m.r.FindByName(username)
}

func (m UserManager) Login(username, claimedPassword string) (string, error) {
	hashedPassword, err := m.r.FindPasswordByName(username)
	if err != nil {
		return "", err
	}

	err = utils.CheckPasswordHash(claimedPassword, hashedPassword)

	if err != nil {
		return "", err
	}

	return jwt.GenerateToken(username)
}
