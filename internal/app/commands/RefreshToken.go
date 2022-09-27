package commands

import "github.com/pvtalbot/money/domain/models"

type RefreshToken struct {
	UserId       string
	RefreshToken string
}

type RefreshTokenHandler struct {
	r models.TokenRepository
	u models.UserRepository
}

func NewRefreshTokenHandler(r models.TokenRepository, u models.UserRepository) RefreshTokenHandler {
	return RefreshTokenHandler{
		u: u,
		r: r,
	}
}

func (h RefreshTokenHandler) Handle(cmd RefreshToken) (*models.Token, error) {
	_, err := h.r.Find(cmd.UserId)
	if err != nil {
		return nil, err
	}

	user, err := h.u.Find(cmd.UserId)
	if err != nil {
		return nil, err
	}

	return h.r.Create(user.ID, user.Name)
}
