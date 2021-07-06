package observer

import (
	"sync"
)

type Event struct {
	Message string
}

type Bus interface {
	Subscribe(topic string, ch <-chan Event)
	Publish(topic string, ch chan<- Event)
}

type AsyncEventBus struct {
	store map[string]chan Event
	mu    sync.Mutex
}

func (a *AsyncEventBus) Subscribe(topic string, ch <-chan Event) {
	panic("implement me")
}

func (a *AsyncEventBus) Publish(topic string, ch chan<- Event) {
	panic("implement me")
}

func NewAsyncEventBus() Bus {
	return &AsyncEventBus{
		store: map[string]chan Event{},
		mu:    sync.Mutex{},
	}
}
