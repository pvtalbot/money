package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pvtalbot/money/domain/model"
)

var (
	SecretKey = []byte("secret")
)

func GenerateToken(username, id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       id,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*model.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		id := claims["id"].(string)
		user := model.User{Name: username, ID: id}
		return &user, nil
	} else {
		return nil, err
	}
}
