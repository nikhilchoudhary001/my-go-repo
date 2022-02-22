package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool, 5)
	go func() {
		<-ch // since we are closing the channel after 2 sec,
		//  it will know that nothing is to be consumed, so it will not wait anymore and move to next line
		fmt.Print("Goroutine")
	}()
	time.Sleep(2 * time.Second)
	close(ch)
	time.Sleep(2 * time.Second)
	fmt.Print("Main")
}
