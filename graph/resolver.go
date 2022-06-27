package graph

import (
	"sync"

	"github.com/digiz3d/graphgogen/graph/model"
)

type Resolver struct {
	ShowsRepository       map[string]*model.Show
	UsersRepository       map[string]*model.User
	ShowCreationObservers map[string]chan *model.CreateShowPayload
	Mu                    sync.Mutex
}
