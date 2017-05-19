package server

import (
	"github.com/funkygao/rpque/pkg/idgen"
)

// Queue is the implementation of replicated priority queue.
type Queue struct {
	id     idgen.IdGenerator
	apisvr *apiServer
}

func New() *Queue {
	apisvr := newApiServer()
	return &Queue{
		apisvr: apisvr,
	}
}

func (q *Queue) Start() error {
	return nil
}

func (q *Queue) ServeForever() {
}
