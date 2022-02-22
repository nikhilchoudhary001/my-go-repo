package main

import "fmt"

// Type assertion in Go - > Assert into concrete type
// **Here Person is a struct and Personnel is a interface
// **Person has two methods - getName() and toString()
// var p personnel -> interface has only 1 method getName()
// p = person{name : “Nikhil”}
// p.getName()
// **Since p is of personnel type we can’t call toString() like p.toString(), therefore
// we need to assert into concrete type- Person
func main() {
	var i interface{}

	i = []int{1, 3}
	//i.([]int)
	fmt.Println(i.([]int)[1])

	fmt.Println("-- MAKE --")
	a := make([]int, 0)
	aPtr := &a
	fmt.Println("pointer == nil :", *aPtr == nil)
	fmt.Printf("pointer value: %v\n\n", *aPtr)
}
