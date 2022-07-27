package graph

import (
	"back_go/internal/domain/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService    model.UserServiceInterface
	ExpenseService model.ExpenseServiceInterface
}
