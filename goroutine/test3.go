package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool)
	wg.Add(2)
	go even(ch, &wg)
	go odd(ch, &wg)
	ch <- true
	wg.Wait()
}

func odd(ch chan bool, wg *sync.WaitGroup) {
	for i := 1; ; i = i + 2 {
		<-ch
		time.Sleep(1 * time.Second)
		fmt.Println(i)
		ch <- true
	}
	wg.Done()
}

func even(ch chan bool, wg *sync.WaitGroup) {
	for i := 2; ; i = i + 2 {
		<-ch
		time.Sleep(1 * time.Second)
		fmt.Println(i)
		ch <- true
	}
	wg.Done()

}
