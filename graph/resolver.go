package graph

import "github.com/digiz3d/graphgogen/graph/model"

type Resolver struct {
	shows map[string]*model.Show
	users map[string]*model.User
}
