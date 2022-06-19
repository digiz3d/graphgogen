package graph

import "github.com/digiz3d/graphgogen/graph/model"

type Resolver struct {
	ShowsRepository map[string]*model.Show
	UsersRepository map[string]*model.User
}
