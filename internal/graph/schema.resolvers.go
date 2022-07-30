package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/pvtalbot/money/app/middlewares"
	"github.com/pvtalbot/money/graph/generated"
	"github.com/pvtalbot/money/graph/model"
	"github.com/pvtalbot/money/pkg/utils"
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
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	expenseService := r.ExpenseService
	expense := expenseService.Create(input.Amount, input.Date, user)

	return &model.Expense{
		ID:     expense.ID,
		Amount: expense.Amount,
		Date:   expense.GetDate(),
	}, nil
}

// DeleteExpense is the resolver for the deleteExpense field.
func (r *mutationResolver) DeleteExpense(ctx context.Context, input model.DeleteExpenseInput) (*model.Expense, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	expense, err := r.ExpenseService.Delete(input.ID, user.ID)

	if err != nil {
		return nil, err
	}

	return &model.Expense{
		ID:     expense.ID,
		Amount: expense.Amount,
		Date:   expense.GetDate(),
	}, nil
}

// UpdateExpense is the resolver for the updateExpense field.
func (r *mutationResolver) UpdateExpense(ctx context.Context, input model.UpdateExpenseInput) (*model.Expense, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	expense, err := r.ExpenseService.Update(input.ID, user.ID, input.Amount, input.Date)

	if err != nil {
		return nil, err
	}

	return &model.Expense{
		ID:     expense.ID,
		Amount: expense.Amount,
		Date:   expense.GetDate(),
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

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Name:      user.Name,
	}, nil
}

// Expenses is the resolver for the expenses field.
func (r *queryResolver) Expenses(ctx context.Context, input model.GetExpensesInput) ([]*model.Expense, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var expenses []*model.Expense
	for _, e := range r.ExpenseService.GetAllExpensesFromUserBetweenDates(user, input.StartDate, input.EndDate) {
		expenses = append(expenses, &model.Expense{ID: e.ID, Amount: e.Amount, Date: e.GetDate()})
	}

	return expenses, nil
}

// ExpensesSum is the resolver for the expensesSum field.
func (r *queryResolver) ExpensesSum(ctx context.Context, input model.GetExpensesSumInput) ([]*model.ExpenseSum, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var groupBy utils.Duration
	switch input.GroupBy {
	case model.DurationMonth:
		groupBy = utils.MONTH
	}

	var expensesSum []*model.ExpenseSum
	for _, e := range r.ExpenseService.SumAllExpensesFromUserBetweenDates(user, input.StartDate, input.EndDate, groupBy) {
		expensesSum = append(expensesSum, &model.ExpenseSum{Amount: e.Amount, StartDate: e.StartDate, EndDate: e.EndDate})
	}

	return expensesSum, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
