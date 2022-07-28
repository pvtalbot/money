package model

type Expense struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

type ExpenseRepository interface {
	Create(expense *Expense, user *User) *Expense
	GetAllExpensesFromUser(user *User) []*Expense
}

type ExpenseServiceInterface interface {
	Create(amount int, user *User) *Expense
}
