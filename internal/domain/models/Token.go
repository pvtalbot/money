package models

type Token struct {
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenRepository interface {
	Create(userId, userName string) (*Token, error)
	Find(userId string) (*Token, error)
}
