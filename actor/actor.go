package actor

import "sync"

type Actor[T any] struct {
	messages chan T
	handler	func(T)
	wg sync.WaitGroup
}

func New[T any](handler func(T)) *Actor[T] {
	a := &Actor[T]{
		messages: make(chan T),
		handler: handler,
	}
	go func() {
		for m := range a.messages {
			a.handler(m)
		}
	}()
	return a
}

func (a *Actor[T]) Send(m T) {
	a.wg.Add(1)
	go func() {
		a.messages <- m
		a.wg.Done()
	}()
}

func (a *Actor[T]) Stop() {
	go func() {
		a.wg.Wait()
		close(a.messages)
	}()
}
