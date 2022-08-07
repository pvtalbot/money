package managers

import (
	"github.com/pvtalbot/money/app/domain/model"
	"github.com/pvtalbot/money/pkg/jwt"
	"github.com/pvtalbot/money/pkg/utils"
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

func (m UserManager) ValidateToken(token string) bool {
	_, err := jwt.ParseToken(token)

	return err == nil
}
