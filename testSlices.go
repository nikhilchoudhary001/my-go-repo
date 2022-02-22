package main

import "fmt"

func main() {
	//s := make([]int, 3, 5)
	s := []int{}
	//s[0] = 1 // panic: runtime error: index out of range [0] with length 0
	//var s []int
	//s1 := []int{}
	//var s [6]int
	fmt.Print(s)
	fmt.Printf("Intial cap %v, len %v, %p\n", cap(s), len(s), s)
	for i := 0; i < 11; i++ {
		s = append(s, i)
		//s[i] = 1
		fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
	}

	p := []int{1, 2, 3, 4, 5, 6, 7, 8}

	p = append(p[:3], p[4:]...)

	fmt.Printf("%v", p)

	a := []string{"A", "B", "C", "D", "E"}
	a = append(a[:2], a[3:]...)
	fmt.Printf("%v", a)

	b := []string{"A", "B", "C", "D", "E"}
	i := 2
	copy(b[i:], b[i+1:]) // Shift a[i+1:] left one index.
	b[len(b)-1] = ""     // Erase last element (write zero value).
	b = b[:len(b)-1]
	fmt.Printf("%v", b)
}
