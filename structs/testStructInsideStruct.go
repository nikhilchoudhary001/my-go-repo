package main

import "fmt"

type Config struct {
	Logger A
	Server B
}

type A struct {
	ValueOfA string
}

func (a *A) func1() string {
	fmt.Println("Inside func1")
	return "fromFunc1"
}

type B struct {
	ValueOfB *string
}

func main() {
	// a := &A{"valueofa"}
	// c := &Config{Logger: a}

	c := &Config{}

	fmt.Println("Test 1 is ", c.Logger.ValueOfA)
	fmt.Println("Test 2 is ", c.Logger.func1())

	fmt.Println("Test 3 is ", c.Server.ValueOfB)
}
