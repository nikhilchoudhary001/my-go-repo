package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0, 10)
	b := a[8:10]
	fmt.Println(len(a), cap(a), len(b), cap(b))
	fmt.Println(b)
	fmt.Println(a)

	a = a[8:10]
	fmt.Println(len(a), cap(a), len(b), cap(b))
	fmt.Println(b)
	fmt.Println(a)

	a = a[0:10]
	fmt.Println(len(a), cap(a), len(b), cap(b))
	fmt.Println(b)
	fmt.Println(a)
}
