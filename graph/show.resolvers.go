package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/digiz3d/graphgogen/graph/generated"
	"github.com/digiz3d/graphgogen/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateShow(ctx context.Context, input model.CreateShowInput) (*model.CreateShowPayload, error) {
	if r.users == nil {
		r.users = make(map[string]*model.User)
	}

	foundUser := r.users[input.UserID]

	if foundUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	show := &model.Show{ID: uuid.NewString(), Name: input.Name, Description: input.Description, UserID: foundUser.ID}

	if r.shows == nil {
		r.shows = make(map[string]*model.Show)
	}
	r.shows[show.ID] = show
	return &model.CreateShowPayload{Show: show}, nil
}

func (r *queryResolver) Show(ctx context.Context, id string) (*model.Show, error) {
	show := r.shows[id]
	if show == nil {
		return nil, fmt.Errorf("show not found")
	}
	return show, nil
}

func (r *showResolver) User(ctx context.Context, obj *model.Show) (*model.User, error) {
	user := r.users[obj.UserID]
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

// Show returns generated.ShowResolver implementation.
func (r *Resolver) Show() generated.ShowResolver { return &showResolver{r} }

type showResolver struct{ *Resolver }
