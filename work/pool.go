package work

import (
	"errors"
	"sync"
)

type Executor interface {
	Execute() error
	OnError(error)
}

type pool struct {
	numWorkers int
	tasks      chan Executor
	start      sync.Once
	stop       sync.Once
	quit       chan struct{}
}

func NewPool(numWorkers int, taskChannelSize int) (*pool, error) {
	if numWorkers <= 0 {
		return nil, errors.New("num workers cannot be less, or equal to 0")
	}

	return &pool{
		numWorkers: numWorkers,
	}, nil
}
