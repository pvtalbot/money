package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pvtalbot/money/domain/models"
)

var (
	secretKey = []byte("secret")
)

func GenerateToken(username, id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       id,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		id := claims["id"].(string)
		user := models.User{Name: username, ID: id}
		return &user, nil
	} else {
		return nil, err
	}
}
