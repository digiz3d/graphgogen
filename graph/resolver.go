package graph

//go:generate go run github.com/99designs/gqlgen generate

import "github.com/digiz3d/graphgogen/graph/model"

type Resolver struct {
	shows []*model.Show
	users []*model.User
}
