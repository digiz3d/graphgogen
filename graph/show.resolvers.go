package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/digiz3d/graphgogen/graph/generated"
	"github.com/digiz3d/graphgogen/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateShow(ctx context.Context, input model.CreateShowInput) (*model.CreateShowPayload, error) {
	foundUser := r.UsersRepository[input.UserID]

	if foundUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	show := &model.Show{ID: uuid.NewString(), Name: input.Name, Description: input.Description, UserID: foundUser.ID}

	r.ShowsRepository[show.ID] = show
	createShowPayload := &model.CreateShowPayload{Show: show}

	r.Mu.Lock()
	for _, observer := range r.ShowCreationObservers {
		observer <- createShowPayload
	}
	r.Mu.Unlock()

	return createShowPayload, nil
}

func (r *queryResolver) Show(ctx context.Context, id string) (*model.Show, error) {
	show := r.ShowsRepository[id]
	if show == nil {
		return nil, fmt.Errorf("show not found")
	}
	return show, nil
}

func (r *showResolver) User(ctx context.Context, obj *model.Show) (*model.User, error) {
	user := r.UsersRepository[obj.UserID]
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *subscriptionResolver) OnCreateShow(ctx context.Context) (<-chan *model.CreateShowPayload, error) {
	id := "thesubkiri" + strconv.Itoa(rand.Int())

	channel := make(chan *model.CreateShowPayload, 1)

	go func() {
		<-ctx.Done()
		r.Mu.Lock()
		delete(r.ShowCreationObservers, id)
		r.Mu.Unlock()
	}()

	r.Mu.Lock()
	r.ShowCreationObservers[id] = channel
	r.Mu.Unlock()

	return channel, nil
}

// Show returns generated.ShowResolver implementation.
func (r *Resolver) Show() generated.ShowResolver { return &showResolver{r} }

type showResolver struct{ *Resolver }
