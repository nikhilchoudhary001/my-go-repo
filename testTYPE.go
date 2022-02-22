package main

import "fmt"

type sample func(u float32)

func (s sample) test(i int) int {
	fmt.Println(i)
	s(65.0)
	return i * i
}

type ex int

func main() {
	v := sample(func(u float32) {
		fmt.Println("Inside sample", u)
	})

	// sample := func(u float32) {
	// 	fmt.Println("Inside sample", u)
	// }
	fmt.Println(v.test(6))
	//fmt.Print(v)
	//v(6.5)
	//sample(7.5)

	x := ex(8)
	fmt.Println(x)
}
