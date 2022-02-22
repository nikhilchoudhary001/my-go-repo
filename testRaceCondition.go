package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 200
	//c := []int{}
	c := make([]int, 200, 200)
	wg := sync.WaitGroup{}
	//count := 0
	//ch := make(chan int, 1)
	wg.Add(n)
	for i := 0; i < n; i++ {
		//fmt.Printf("%p \n", c)
		i := i
		go func() {
			//c = append(c, 1)
			c[i] = 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(len(c)) //?
}
