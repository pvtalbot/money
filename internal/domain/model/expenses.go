package model

import "time"

type Expense struct {
	ID     string    `json:"id"`
	Amount int       `json:"amount"`
	Date   time.Time `json:"date"`
	User   User
}

type ExpenseSum struct {
	Amount    int       `json:"amount"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"EndDate"`
}

type ExpenseRepository interface {
	Create(expense *Expense, user *User) *Expense
	Delete(ID int64) error
	Find(ID int64) (*Expense, error)
	GetAllExpensesFromUserBetweenDates(user *User, startDate, endDate time.Time) []*Expense
	SumAllExpensesFromUserBetweenDatesByMonth(user *User, startDate, endDate time.Time) []*ExpenseSum
}

type ExpenseServiceInterface interface {
	Create(amount int, date time.Time, user *User) *Expense
	GetAllExpensesFromUserBetweenDates(user *User, startDate, endDate time.Time) []*Expense
	Delete(id string, userID string) (*Expense, error)
	SumAllExpensesFromUserBetweenDates(user *User, startDate, endDate time.Time) []*ExpenseSum
}
