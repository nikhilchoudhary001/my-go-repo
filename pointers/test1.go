package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
}

// https://stackoverflow.com/questions/26159416/init-array-of-structs-in-go
func main() {
	p := []*person{
		{
			name: "John1eeffef",
		},
		{
			name: "John2",
		},
		{
			name: "John3",
		},
	}

	fmt.Printf("%p \n", p[0])

	for i, pp := range p {
		fmt.Printf("%p \n", pp)
		pp.name = "Updating Jhon" + strconv.Itoa(i)
	}
	for _, pp := range p {
		fmt.Println(pp.name)
	}
}
