package main

import (
	"fmt"
	"sync"
)

func main() {

	chEven := make(chan int)
	chOdd := make(chan int)

	go even(chEven)
	go odd(chOdd)

	var wg sync.WaitGroup

	wg.Add(1)

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			chOdd <- i
		} else {
			chEven <- i

		}
	}
	wg.Done()
	wg.Wait()

}
func odd(ch chan int) {
	for v := range ch {
		fmt.Println("Goroutine 1 produced: odd number%d", v)
	}
}

func even(ch chan int) {
	for v := range ch {
		fmt.Println("Goroutine 2 produced: Even number%d", v)
	}

}
