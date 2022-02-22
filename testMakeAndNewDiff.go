package main

import "fmt"

func main() {

	var s []int
	//s := []int{}
	//s[10] = 20
	fmt.Println(s)
	fmt.Println("-- MAKE --")
	a := make([]int, 0)
	aPtr := &a
	fmt.Println("pointer == nil :", *aPtr == nil)
	fmt.Printf("pointer value: %p\n\n", *aPtr)

	fmt.Println("-- COMPOSITE LITERAL --")
	b := []int{}
	bPtr := &b
	fmt.Println(bPtr)
	fmt.Println("pointer == nil :", *bPtr == nil)
	fmt.Printf("pointer value: %p\n\n", *bPtr)

	fmt.Println("-- NEW --")
	cPtr := new([]int)
	fmt.Println("pointer == nil :", *cPtr == nil)
	fmt.Printf("pointer value: %p\n\n", *cPtr)

	fmt.Println("-- VAR (not initialized) --")
	var d []int
	dPtr := &d
	fmt.Println(dPtr)
	fmt.Println("pointer == nil :", *dPtr == nil)
	fmt.Printf("pointer value: %p\n", *dPtr)

	// NEW and VAR (not initialized)  bahave same way
}
