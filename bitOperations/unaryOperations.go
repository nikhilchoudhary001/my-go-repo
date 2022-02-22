package main

import (
	"fmt"
)

func main() {
	a := 7
	b := 11

	// c := "4.7.91.208"
	// cidr := "4.7.95.208/29"

	// res := strings.Split(cidr, "/")

	// AND operations
	fmt.Println(a & b)

	// Left shift operation
	j := 62
	fmt.Println(1 << j)

	// Right shift operation
	r := 5
	fmt.Println(1 >> r) // will always give 0 as answer
}
