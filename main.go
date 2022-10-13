package main

import (
	"fmt"
	"os"
)

func main() {
    arg := os.Args[1]

    if arg == "helloworld" {
        fmt.Print("Hello, World!")
        helloworld()
    }
}