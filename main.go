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
	}
}
