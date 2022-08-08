package model

type User struct {
	ID                 string            `json:"id"`
	Name               string            `json:"name"`
	FirstName          string            `json:"firstName"`
	LastName           string            `json:"lastName"`
	ExpensesCategories []ExpenseCategory `json:"expensesCategories"`
}
