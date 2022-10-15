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
	}
}
