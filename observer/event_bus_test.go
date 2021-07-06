package observer

import (
	"fmt"
	"testing"
	"time"
)

func TestEventBus(t *testing.T) {
	bus := NewAsyncEventBus()

	bus.Publish("/hello", Event{Message: "hello world"})
	_ = bus.Subscribe("/hello", func(events ...Event) {
		for _, e := range events {
			fmt.Println(e.Message)
		}
	})
	bus.Publish("/hello", Event{Message: "aaa"})


	time.Sleep(time.Second)
}