package commands

import (
	"github.com/pvtalbot/money/domain/models"
	"github.com/pvtalbot/money/pkg/utils"
)

type Login struct {
	Name            string
	ClaimedPassword string
}

type LoginHandler struct {
	r               models.UserRepository
	tokenRepository models.TokenRepository
}

func NewLoginHandler(r models.UserRepository, tokenRepository models.TokenRepository) LoginHandler {
	return LoginHandler{
		r:               r,
		tokenRepository: tokenRepository,
	}
}

func (h LoginHandler) Handle(cmd Login) (*models.Token, error) {
	hashedPassword, err := h.r.FindPasswordByName(cmd.Name)

	if err != nil {
		return nil, err
	}

	err = utils.CheckPasswordHash(cmd.ClaimedPassword, hashedPassword)
	if err != nil {
		return nil, err
	}

	user, err := h.r.FindByName(cmd.Name)
	if err != nil {
		return nil, err
	}

	return h.tokenRepository.Create(user.ID, user.Name)
}
