package model

type Expense struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

type ExpenseRepository interface {
	Create(expense *Expense) *Expense
}

type ExpenseServiceInterface interface {
	Create(amount int) *Expense
}
