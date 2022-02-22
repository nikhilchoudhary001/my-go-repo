package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:4]
	t[1] = 9
	fmt.Println(t)

	// Remove index no. 2 element from the slice b1
	b1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Len and cap ", len(b1), cap(b1))
	b1 = append(b1, 6)
	// Although we increased the capacity, but we can't do b1[6] as length is still 6 and we are trying to access 7th index
	b1[5] = 7
	fmt.Println("Len and cap ", len(b1), cap(b1))
	b1 = append(b1[0:2], b1[3:10]...)

	fmt.Println("b1 is : ", b1)

	in := 10
	x := &in
	fmt.Println(*x)

	//var b []int
	var a1 [4]int
	var b [4]int
	b = a1
	a1[2] = 5
	fmt.Println(a1, b)

	c := make([]int, 4)
	d := c
	c[2] = 1
	fmt.Println(c, d)

	e := c[1:4]
	e[1] = 20
	fmt.Println(c, d, e)

	var k []int
	fmt.Println(k)
}
