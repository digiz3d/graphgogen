package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/digiz3d/graphgogen/graph/generated"
	"github.com/digiz3d/graphgogen/graph/model"
)

func (r *queryResolver) Show(ctx context.Context, id string) (*model.Show, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *showResolver) User(ctx context.Context, obj *model.Show) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Show returns generated.ShowResolver implementation.
func (r *Resolver) Show() generated.ShowResolver { return &showResolver{r} }

type showResolver struct{ *Resolver }
