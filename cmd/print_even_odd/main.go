package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello even odd")

	wg := &sync.WaitGroup{}
	ch := make(chan struct{}, 1)
	wg.Add(2)

	// ch <- struct{}{}
	go printEven(wg, ch)
	go printOdd(wg, ch)

	wg.Wait()

	fmt.Println("Processing completed")
}

func printEven(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()

	for i := 2; i <= 100; i = i + 2 { // 4
		// if i == 2 {
		// 	<-ch
		// 	fmt.Println(i) // 2
		// } else {
		// 	ch <- struct{}{}
		// 	<-ch
		// 	fmt.Println(i) // 98
		// }
		if i != 2 {
			ch <- struct{}{}
		}
		<-ch
		fmt.Println(i) // 98
	}
}

func printOdd(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()

	for i := 1; i < 100; i = i + 2 { // 3
		// if i == 1 {
		// 	fmt.Println(i) // 1
		// 	ch <- struct{}{}
		// } else {
		// 	<-ch           // waiting
		// 	fmt.Println(i) // 99
		// 	ch <- struct{}{}
		// }
		if i != 1 {
			<-ch
		}
		fmt.Println(i) // 99
		ch <- struct{}{}
	}
}

// 1 2
