package models

import (
	"github.com/pvtalbot/money/pkg/utils"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Expenses  []Expense `json:"expenses"`
	password  string
}

type UserRepository interface {
	Create(user *User) (*User, error)
	Find(id string) (*User, error)
	FindByName(username string) (*User, error)
	FindPasswordByName(username string) (string, error)
}

func (u *User) SetPassword(clearPassword string) (*User, error) {
	hashedPassword, err := utils.HashPassword(clearPassword)
	if err != nil {
		return u, err
	}

	u.password = hashedPassword
	return u, nil
}

func (u *User) GetHashedPassword() string {
	return u.password
}
