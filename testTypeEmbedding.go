package main

import (
	"fmt"
)

type Name struct {
	string
	int
}

func main() {

	// s := Name{"hello", 1}
	var s Name
	s.string = "Nikhil"
	fmt.Println("s", s.string)
	fmt.Println("s1", s.int)

	s.string = "ValueChanged"
	fmt.Println("new s gets mutated", s.string)
}
