package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan string)
	go recive(&wg, ch)
	go send(ch)
	wg.Wait()
	//time.Sleep(10 * time.Second)
	fmt.Println("end of main")
}

func recive(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	data := []string{"1", "2", "3", "4"}
	for i := 0; i < len(data); i++ {
		fmt.Println("Produced ", data[i])
		ch <- data[i]
	}
	close(ch)
}

func send(ch chan string) {
	//time.Sleep(1 * time.Second)
	for {
		fmt.Println("Consumed ", <-ch)
	}

	// for data := range ch {
	// 	fmt.Println("Consumed ", data)
	// }
}
