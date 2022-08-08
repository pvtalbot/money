package models

type ExpenseCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	User *User  `json:"user"`
}

type ExpenseCategoryRepository interface {
	Create(expenseCategory *ExpenseCategory) (*ExpenseCategory, error)
}
