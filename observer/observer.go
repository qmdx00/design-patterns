package observer

import "fmt"

// ISubject 发布者 ...
type ISubject interface {
	Register(observers ...IObserver)
	Remove(observers ...IObserver)
	Notify(msg string)
}

// IObserver 观察者/订阅者 ...
type IObserver interface {
	Update(msg string)
	Key() string
}

// 发布者实现
type Subject struct {
	observers map[string]IObserver
}

func NewSubject() ISubject {
	return &Subject{
		observers: make(map[string]IObserver),
	}
}

func (s *Subject) Register(observers ...IObserver) {
	for _, observer := range observers {
		s.observers[observer.Key()] = observer
	}
}

func (s *Subject) Remove(observers ...IObserver) {
	for _, observer := range observers {
		delete(s.observers, observer.Key())
	}
}

func (s *Subject) Notify(msg string) {
	for _, observer := range s.observers {
		observer.Update(msg)
	}
}

// 观察者 A 实现
type A struct{}

func (a A) Key() string {
	return "a"
}

func (a A) Update(msg string) {
	fmt.Println("observer A:", msg)
}

// 观察者 B 实现
type B struct{}

func (b B) Key() string {
	return "b"
}

func (b B) Update(msg string) {
	fmt.Println("observer B:", msg)
}
