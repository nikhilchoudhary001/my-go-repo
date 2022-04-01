package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool)
	wg.Add(2)
	go even(ch, &wg)
	go odd(ch, &wg)
	wg.Wait()
}

func odd(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Print(i, " Odd ")
		if i == 4 {
			return
		}
	}
}

func even(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 6; i++ {
		fmt.Print(i, " Even ")

	}
}
