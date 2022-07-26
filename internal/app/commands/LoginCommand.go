package commands

import (
	"github.com/pvtalbot/money/domain/models"
	"github.com/pvtalbot/money/pkg/jwt"
	"github.com/pvtalbot/money/pkg/utils"
)

type Login struct {
	Name            string
	ClaimedPassword string
}

type LoginHandler struct {
	r models.UserRepository
}

func NewLoginHandler(r models.UserRepository) LoginHandler {
	return LoginHandler{
		r: r,
	}
}

func (h LoginHandler) Handle(cmd Login) (string, error) {
	hashedPassword, err := h.r.FindPasswordByName(cmd.Name)

	if err != nil {
		return "", err
	}

	err = utils.CheckPasswordHash(cmd.ClaimedPassword, hashedPassword)
	if err != nil {
		return "", err
	}

	user, err := h.r.FindByName(cmd.Name)
	if err != nil {
		return "", err
	}

	return jwt.GenerateToken(user.Name, user.ID)
}
