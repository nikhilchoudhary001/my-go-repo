package main

import (
	"fmt"
)

// Defer executes even if there is a panic occurred in function
func main() {
	// defer fmt.Println("main received", 1)
	// defer fmt.Println("main received", 2)
	// defer fmt.Println("main received", 3)

	// panic("PANIC")

	defer_call()
}

//similar to above code
func defer_call() {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()

	panic("PANIC")
}
