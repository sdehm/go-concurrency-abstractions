package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "helloworld":
		helloworld()
	case "task":
		t := New(func() {
			fmt.Println("Hello, World!")
		})
		t.Start()
		t.Wait()
	case "task_input":
		greeting := "Hello from the closure!"
		t := New(func() {
			fmt.Println(greeting)
		})
		t.Start()
		t.Wait()
	case "task_output":
		var greeting string
		t := New(func() {
			greeting = "Hello from the closure!"
		})
		t.Start()
		t.Wait()
		fmt.Println(greeting)
	case "task_input_generic":
		t := NewWithInput(func(i string) {
			fmt.Println(i)
		}, "Hello with generics!")
		t.Start()
		t.Wait()
	case "task_output_generic":
		t := NewWithResult(func() string {
			return "Hello with generic output!"
		})
		t.Start()
		fmt.Println(t.GetResult())
    case "task_nested":
        t := New(func() {
            fmt.Println("Hello from the first task!")
            t := New(func() {
                fmt.Println("Hello from the second task!")
            })
            t.Start()
            t.Wait()
        })
        t.Start()
        t.Wait()
	}
}
