package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/digiz3d/graphgogen/graph/generated"
	"github.com/digiz3d/graphgogen/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, username string) (*model.User, error) {
	newUser := &model.User{
		ID:       "1",
		Username: "ok",
		Shows:    []*model.Show{},
	}
	if r.users == nil {
		r.users = make(map[string]*model.User)
	}

	r.users[newUser.ID] = newUser
	return newUser, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user := r.users[id]
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *userResolver) Shows(ctx context.Context, obj *model.User) ([]*model.Show, error) {
	return []*model.Show{}, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
