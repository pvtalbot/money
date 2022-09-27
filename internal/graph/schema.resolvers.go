package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/pvtalbot/money/app/commands"
	"github.com/pvtalbot/money/app/queries"
	"github.com/pvtalbot/money/domain/managers"
	custom_errors "github.com/pvtalbot/money/errors"
	"github.com/pvtalbot/money/graph/generated"
	"github.com/pvtalbot/money/graph/model"
	"github.com/pvtalbot/money/middlewares"
	"github.com/pvtalbot/money/pkg/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.Token, error) {
	token, err := r.Application.Commands.Login.Handle(commands.Login{
		Name:            input.Username,
		ClaimedPassword: input.Password,
	})

	if err != nil {
		return nil, err
	}

	return &model.Token{
		AuthToken:    &token.AuthToken,
		RefreshToken: &token.RefreshToken,
	}, nil
}

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.Token, error) {
	cmd := commands.NewCreateUser(
		input.Name,
		input.Password,
		input.FirstName,
		input.LastName,
	)

	token, err := r.Application.Commands.CreateUser.Handle(cmd)
	if err != nil {
		switch err.(type) {
		case custom_errors.DuplicateEntityError:
			responseError := gqlerror.Errorf("username already exists")
			responseError.Extensions = map[string]interface{}{"code": "createuser-1"}
			graphql.AddError(ctx, responseError)
		default:
			graphql.AddError(ctx, gqlerror.Errorf("internal error"))
		}
	}

	return &model.Token{
		AuthToken:    &token.AuthToken,
		RefreshToken: &token.RefreshToken,
	}, nil
}

// CreateExpense is the resolver for the createExpense field.
func (r *mutationResolver) CreateExpense(ctx context.Context, input model.CreateExpenseInput) (*model.Expense, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	category, err := r.Application.Queries.FindExpenseCategory.Handle(queries.FindExpenseCategory{Id: input.CategoryID})
	if err != nil {
		return nil, err
	}
	if category.User.ID != user.ID {
		return nil, errors.New("this category does not belong to the current user")
	}

	cmd := commands.CreateExpense{
		Amount:     input.Amount,
		Date:       input.Date,
		UserId:     user.ID,
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

	if input.CategoryID != nil {
		category, err := r.Application.Queries.FindExpenseCategory.Handle(queries.FindExpenseCategory{Id: *input.CategoryID})
		if err != nil {
			return nil, err
		}
		if category.User.ID != user.ID {
			return nil, errors.New("this category does not belong to the current user")
		}
	}

	cmd := commands.NewUpdateExpenseCommand(*expense, input.Amount, input.Date, input.CategoryID)
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

// CreateRevenue is the resolver for the createRevenue field.
func (r *mutationResolver) CreateRevenue(ctx context.Context, input model.CreateRevenueInput) (*model.Revenue, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	cmd := commands.CreateRevenue{
		Amount: input.Amount,
		Date:   input.Date,
		UserId: user.ID,
	}

	revenue, err := r.Application.Commands.CreateRevenue.Handle(cmd)
	if err != nil {
		return nil, err
	}

	return &model.Revenue{
		ID:     revenue.ID,
		Amount: revenue.Amount,
		Date:   revenue.GetDate(),
	}, nil
}

// DeleteRevenue is the resolver for the deleteRevenue field.
func (r *mutationResolver) DeleteRevenue(ctx context.Context, input model.DeleteRevenueInput) (*model.Revenue, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	revenue, err := r.Application.Queries.FindRevenue.Handle(
		queries.FindRevenue{Id: input.ID},
	)
	if err != nil {
		return nil, err
	}

	if user.ID != revenue.User.ID {
		return nil, errors.New("user cannot delete revenue")
	}

	err = r.Application.Commands.DeleteRevenue.Handle(
		commands.DeleteRevenue{Id: revenue.ID},
	)
	if err != nil {
		return nil, err
	}

	return &model.Revenue{
		ID:     revenue.ID,
		Amount: revenue.Amount,
		Date:   revenue.GetDate(),
	}, nil
}

// UpdateRevenue is the resolver for the updateRevenue field.
func (r *mutationResolver) UpdateRevenue(ctx context.Context, input model.UpdateRevenueInput) (*model.Revenue, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	revenue, err := r.Application.Queries.FindRevenue.Handle(
		queries.FindRevenue{Id: input.ID},
	)
	if err != nil {
		return nil, err
	}
	if user.ID != revenue.User.ID {
		return nil, errors.New("user cannot update revenue")
	}

	log.Println("kikoo")

	cmd := commands.NewUpdateRevenueCommand(*revenue, input.Amount, input.Date)
	revenue, err = r.Application.Commands.UpdateRevenue.Handle(cmd)

	log.Println("kikoo 2")

	if err != nil {
		return nil, err
	}

	return &model.Revenue{
		ID:     revenue.ID,
		Amount: revenue.Amount,
		Date:   revenue.GetDate(),
	}, nil
}

// GetAllErrors is the resolver for the getAllErrors field.
func (r *queryResolver) GetAllErrors(ctx context.Context) ([]*model.Error, error) {
	query := queries.GetAllErrors{}
	errors, _ := r.Application.Queries.GetAllErrors.Handle(query)

	var result []*model.Error
	for _, e := range errors {
		result = append(result, &model.Error{Code: e.Code, Translation: e.Translation})
	}

	return result, nil
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

	query := queries.NewGetExpensesQuery(user.ID, input.StartDate, input.EndDate)
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

// Revenues is the resolver for the revenues field.
func (r *queryResolver) Revenues(ctx context.Context, input model.GetRevenuesInput) ([]*model.Revenue, error) {
	user, err := middlewares.ExtractUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	query := queries.NewGetRevenuesQuery(user.ID, input.StartDate, input.EndDate)
	rev, err := r.Application.Queries.GetRevenues.Handle(query)

	if err != nil {
		return nil, err
	}

	var revenues []*model.Revenue
	for _, r := range rev {
		revenues = append(revenues, &model.Revenue{
			ID:     r.ID,
			Amount: r.Amount,
			Date:   r.GetDate(),
		})
	}

	return revenues, nil
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
		user.ID,
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

// ValidateAuthToken is the resolver for the validateAuthToken field.
func (r *queryResolver) ValidateAuthToken(ctx context.Context, authToken string) (bool, error) {
	return managers.ValidateToken(authToken), nil
}

// ExpensesCategories is the resolver for the expensesCategories field.
func (r *userResolver) ExpensesCategories(ctx context.Context, obj *model.User) ([]*model.ExpenseCategory, error) {
	var expensesCategories []*model.ExpenseCategory

	result, err := r.Application.Queries.GetExpensesCategories.Handle(queries.GetExpensesCategories{
		UserId: obj.ID,
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
