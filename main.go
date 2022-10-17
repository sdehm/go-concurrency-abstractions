package main

import (
	"fmt"
	"os"

	"github.com/sdehm/go-concurrency-abstractions/task"
	"github.com/sdehm/go-concurrency-abstractions/workers"
)

func main() {
	arg := os.Args[1]

	switch arg {
	// Helloworld example
	case "helloworld":
		helloworld()
		// Task examples
	case "task":
		t := task.New(func() {
			fmt.Println("Hello, World!")
		})
		t.Start()
		t.Wait()
	case "task_input":
		greeting := "Hello from the closure!"
		t := task.New(func() {
			fmt.Println(greeting)
		})
		t.Start()
		t.Wait()
	case "task_output":
		var greeting string
		t := task.New(func() {
			greeting = "Hello from the closure!"
		})
		t.Start()
		t.Wait()
		fmt.Println(greeting)
	case "task_input_generic":
		t := task.NewWithInput(func(i string) {
			fmt.Println(i)
		}, "Hello with generics!")
		t.Start()
		t.Wait()
	case "task_output_generic":
		t := task.NewWithResult(func() string {
			return "Hello with generic output!"
		})
		t.Start()
		fmt.Println(t.GetResult())
	case "task_nested":
		t := task.New(func() {
			fmt.Println("Hello from the first task!")
			t := task.New(func() {
				fmt.Println("Hello from the second task!")
			})
			t.Start()
			t.Wait()
		})
		t.Start()
		t.Wait()
	// Worker examples
	case "worker":
		w := workers.New[string](2)
		go func() {
			for i := 0; i < 10; i++ {
				i := i
				w.Work <- func() string {
					return fmt.Sprintf("%d", i)
				}
			}
			close(w.Work)
		}()
		for r := range w.Results {
			fmt.Println(r)
		}
	}
}
