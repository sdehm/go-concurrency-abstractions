package main

// This type implements a task abstraction loosely based on .NET Tasks.
type Task struct {
	f       func()
	awaiter chan struct{}
}

// Creates a new task.
func New(f func()) *Task {
	return &Task{
		f:       f,
		awaiter: make(chan struct{}),
	}
}

// Executes the task.
func (t *Task) Start() {
	go func() {
		t.f()
		t.awaiter <- struct{}{}
	}()
}

// Waits for the task to complete.
func (t *Task) Wait() {
	<-t.awaiter
}
