package main

import "fmt"

type MyError struct{}

func (m *MyError) Error() string {
	return "My Error"
}

func f() *MyError { return nil }

func main() {

	var err error
	err = f() // Value will be nil but the type is not nil

	if err != nil {
		fmt.Println("something went wrong")
	} else {
		fmt.Println("nothing to see here")
	}
}
