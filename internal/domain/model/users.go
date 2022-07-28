package model

import (
	"back_go/pkg/utils"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Expenses  []Expense `json:"expenses"`
	password  string
}

type UserServiceInterface interface {
	Create(username, password, firstName, LastName string) *User
	FindAll() []*User
	FindByName(username string) (*User, error)
	Login(username, claimedPassword string) (string, error)
	GetAllExpensesFromUser(user *User) []*Expense
}

type UserRepository interface {
	Create(user *User) *User
	FindAll() []*User
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
