package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/pvtalbot/money/app/commands"
	"github.com/pvtalbot/money/app/queries"
	"github.com/pvtalbot/money/domain/managers"
	"github.com/pvtalbot/money/domain/models"
	"github.com/pvtalbot/money/graph/generated"
	"github.com/pvtalbot/money/graph/model"
	"github.com/pvtalbot/money/middlewares"
	"github.com/pvtalbot/money/pkg/utils"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return r.Application.Commands.Login.Handle(commands.Login{
		Name:            input.Username,
		ClaimedPassword: input.Password,
	})
}

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	cmd := commands.NewCreateUser(
		input.Name,
		input.Password,
		input.FirstName,
		input.LastName,
	)

	user, err := r.Application.Commands.CreateUser.Handle(cmd)
	if err != nil {
		return nil, err
	}

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

	category, err := r.Application.Queries.FindExpenseCategory.Handle(queries.FindExpense{Id: input.CategoryID})
	if err != nil {
		return nil, err
	}
	if category.User.ID != user.ID {
		return nil, errors.New("this category does not belong to the current user")
	}

	cmd := commands.CreateExpense{
		Amount:     input.Amount,
		Date:       input.Date,
		User:       user,
		CategoryId: input.CategoryID,
	}

	expense, err := r.Application.Commands.CreateExpense.Handle(cmd)
	if err != nil {
		return nil, err
	}

	return &model.Expense{
		ID:       expense.ID,
		Amount:   expense.Amount,
		Date:     expense.GetDate(),
		Category: &model.ExpenseCategory{ID: input.CategoryID},
	}, nil
}

// DeleteExpense is the resolver for the deleteExpense field.
func (r *mutationResolver) DeleteExpense(ctx context.Context, input model.DeleteExpenseInput) (*model.Expense, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	expense, err := r.Application.Queries.FindExpense.Handle(
		queries.FindExpense{Id: input.ID},
	)
	if err != nil {
		return nil, err
	}

	if user.ID != expense.User.ID {
		return nil, errors.New("user cannot delete expense")
	}

	err = r.Application.Commands.DeleteExpense.Handle(
		commands.DeleteExpense{Id: expense.ID},
	)

	if err != nil {
		return nil, err
	}

	return &model.Expense{
		ID:       expense.ID,
		Amount:   expense.Amount,
		Date:     expense.GetDate(),
		Category: &model.ExpenseCategory{ID: expense.Category.ID},
	}, nil
}

// UpdateExpense is the resolver for the updateExpense field.
func (r *mutationResolver) UpdateExpense(ctx context.Context, input model.UpdateExpenseInput) (*model.Expense, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	expense, err := r.Application.Queries.FindExpense.Handle(
		queries.FindExpense{Id: input.ID},
	)
	if err != nil {
		return nil, err
	}
	if user.ID != expense.User.ID {
		return nil, errors.New("user cannot update expense")
	}

	cmd := commands.NewUpdateExpenseCommand(*expense, input.Amount, input.Date)
	expense, err = r.Application.Commands.UpdateExpense.Handle(cmd)

	if err != nil {
		return nil, err
	}

	return &model.Expense{
		ID:       expense.ID,
		Amount:   expense.Amount,
		Date:     expense.GetDate(),
		Category: &model.ExpenseCategory{ID: expense.Category.ID},
	}, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err = r.Application.Queries.FindUser.Handle(
		queries.FindUser{Id: user.ID},
	)
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

	query := queries.NewGetExpensesQuery(user, input.StartDate, input.EndDate)
	exp, err := r.Application.Queries.GetExpenses.Handle(query)

	if err != nil {
		return nil, err
	}

	var expenses []*model.Expense
	for _, e := range exp {
		expenses = append(expenses, &model.Expense{
			ID:       e.ID,
			Amount:   e.Amount,
			Date:     e.GetDate(),
			Category: &model.ExpenseCategory{ID: e.Category.ID},
		})
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
	query := queries.NewSumExpensesQuery(
		input.StartDate,
		input.EndDate,
		user,
		groupBy,
	)
	result, err := r.Application.Queries.SumExpenses.Handle(query)
	if err != nil {
		return nil, err
	}

	for _, e := range result {
		expensesSum = append(expensesSum, &model.ExpenseSum{Amount: e.Amount, StartDate: e.StartDate, EndDate: e.EndDate})
	}

	return expensesSum, nil
}

// ValidateAccessToken is the resolver for the validateAccessToken field.
func (r *queryResolver) ValidateAccessToken(ctx context.Context, accessToken string) (bool, error) {
	return managers.ValidateToken(accessToken), nil
}

// ExpensesCategories is the resolver for the expensesCategories field.
func (r *userResolver) ExpensesCategories(ctx context.Context, obj *model.User) ([]*model.ExpenseCategory, error) {
	var expensesCategories []*model.ExpenseCategory

	result, err := r.Application.Queries.GetExpensesCategories.Handle(queries.GetExpensesCategories{
		User: &models.User{ID: obj.ID},
	})
	if err != nil {
		return nil, err
	}

	for _, v := range result {
		expensesCategories = append(expensesCategories, &model.ExpenseCategory{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return expensesCategories, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
