package managers

import (
	"github.com/pvtalbot/money/pkg/jwt"
)

func ValidateToken(token string) bool {
	_, err := jwt.ParseToken(token)

	return err == nil
}
