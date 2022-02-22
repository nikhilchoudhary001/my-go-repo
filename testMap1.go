package main

import "fmt"

func main() {
	m := make(map[int]int)

	for i := 0; i < 50; i++ {
		m[i] = 50 - i
	}

	for key, v := range m {
		fmt.Println(key, v)
	}
}
