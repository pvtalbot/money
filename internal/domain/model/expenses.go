package model

import (
	"back_go/pkg/utils"
	"time"
)

type Expense struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
	User   User   `json:"user"`
	date   time.Time
}

type ExpenseSum struct {
	Amount    int       `json:"amount"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"EndDate"`
}

func (e *Expense) GetDate() time.Time {
	return e.date
}

func (e *Expense) SetDate(d time.Time) {
	/*
		as we care only about the date and not the time of the expense, all expenses are set at 12PM. As we request the table
		with datetimes set at 0AM, it's easy to guarantee we have all relevant expenses
	*/
	roundedTime := time.Date(d.Year(), d.Month(), d.Day(), 12, 0, 0, 0, d.Location())
	e.date = roundedTime
}

type ExpenseRepository interface {
	Create(expense *Expense, user *User) *Expense
	Delete(ID string) error
	Find(ID string) (*Expense, error)
	Update(expense *Expense) (*Expense, error)
	GetAllExpensesFromUserBetweenDates(user *User, startDate, endDate time.Time) []*Expense
	SumAllExpensesFromUserBetweenDatesByMonth(user *User, startDate, endDate time.Time) []*ExpenseSum
}

type ExpenseServiceInterface interface {
	Create(amount int, date time.Time, user *User) *Expense
	Delete(id string, userID string) (*Expense, error)
	Update(id, userId string, amount *int, date *time.Time) (*Expense, error)
	GetAllExpensesFromUserBetweenDates(user *User, startDate, endDate time.Time) []*Expense
	SumAllExpensesFromUserBetweenDates(user *User, startDate, endDate time.Time, groupBy utils.Duration) []*ExpenseSum
}
