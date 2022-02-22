package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7}
	fmt.Println(Sum(primes...)) // 17
}

func Sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
