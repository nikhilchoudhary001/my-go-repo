package main

import "fmt"

// copy function has slice as both args.
func main() {
	a := []int{1, 2}
	b := []int{3, 4}

	check := a
	copy(a, b)
	fmt.Println(a, b, check)
}
