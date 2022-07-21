package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"back_go/graph/generated"
	"back_go/graph/model"
	"back_go/internal/pkg/db/users"
	"context"
	"strconv"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return "kikoo", nil
}

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	var user users.User
	user.Name = input.Name
	user.Password = input.Password
	userId := user.Save()
	return &model.User{
		ID:   strconv.FormatInt(userId, 10),
		Name: user.Name,
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var resultUsers []*model.User
	dbUsers := users.GetAll()
	for _, user := range dbUsers {
		resultUsers = append(resultUsers, &model.User{ID: user.ID, Name: user.Name})
	}

	return resultUsers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
