package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	a := &A{}
	b := &B{}

	s := NewSubject()
	s.Register(a, b)
	t.Log(s)
	s.Notify("hello")

	s.Remove(a)
	t.Log(s)
	s.Notify("world")
}
