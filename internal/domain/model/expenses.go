package model

import "time"

type Expense struct {
	ID     string    `json:"id"`
	Amount int       `json:"amount"`
	Date   time.Time `json:"date"`
}

type ExpenseRepository interface {
	Create(expense *Expense, user *User) *Expense
	GetAllExpensesFromUserBetweenDates(user *User, startDate, endDate time.Time) []*Expense
}

type ExpenseServiceInterface interface {
	Create(amount int, date time.Time, user *User) *Expense
	GetAllExpensesFromUserBetweenDates(user *User, startDate, endDate time.Time) []*Expense
}
