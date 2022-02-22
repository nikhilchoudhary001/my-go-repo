package main

import (
	"fmt"
)

func main() {
	res := foo1()
	fmt.Println(res)

	n := foo()
	fmt.Println("main received", n)
}

func main2() {
	fmt.Println("Hello")
	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("World")
}

func foo1() (result string) {
	defer func() {
		result = "Change World 4" // change value at the very last moment
	}()
	return "Hello World 5"
}

func foo() int {
	m := 6
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occured ", err)
			m = 2 // will not work. and m returned will be 0 only as we are not using (m int) as return type of this func.
		}
	}()
	panic("foo: fail")
	m = 3
	return m
}
