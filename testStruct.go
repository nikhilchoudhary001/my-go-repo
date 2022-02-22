package main

import "fmt"

// type Y struct {
// 	name string
// 	n    X
// }

//next field must be a pointer. It’s not possible to make declaration like:
// type X struct {
// 	name string
// 	next X
// }
//since compiler will respond “invalid recursive type X” because while creating type X and in order to calculate its size compiler finds field next of type X for which it has to do the same —define its size so we’ve an infinite recursion. With pointer it’s doable as compiler knows the size of the pointer for desired architecture.

// func main() {
// 	fmt.Println("main")
// 	funtest()
// }

type Customer struct {
	name string
}

func main() {
	c := new(Customer)
	fmt.Printf("%p \n", *c)
	fmt.Printf("%v \n ", *c)
	c.name = "Nik"
	fmt.Printf("%p \n ", *c)
	fmt.Printf("%v \n", *c)

	g := &Customer{}
	fmt.Printf("%p \n", *g)
	fmt.Printf("%v \n ", *g)

	// Both behave same way

}
