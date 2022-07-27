package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"back_go/graph/generated"
	"back_go/graph/model"
	"context"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return r.UserService.Login(input.Username, input.Password)
}

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	userService := r.UserService
	user := userService.Create(input.Name, input.Password, input.FirstName, input.LastName)

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

// CreateExpense is the resolver for the createExpense field.
func (r *mutationResolver) CreateExpense(ctx context.Context, input model.CreateExpenseInput) (*model.Expense, error) {
	expenseService := r.ExpenseService
	expense := expenseService.Create(input.Amount)

	return &model.Expense{
		ID:     expense.ID,
		Amount: expense.Amount,
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var resultUsers []*model.User
	internalUsers := r.UserService.FindAll()

	for _, user := range internalUsers {
		resultUsers = append(resultUsers, &model.User{ID: user.ID, Name: user.Name, FirstName: user.FirstName, LastName: user.LastName})
	}

	return resultUsers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
