package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)

	go func() {
		//time.Sleep(1 * time.Second)
		c <- true
		fmt.Print("After 1 \n")
		//time.Sleep(1 * time.Second)
		fmt.Print("Before 2 \n")
		c <- false
		fmt.Print("After 2 \n")
		time.Sleep(1 * time.Second)
		c <- true
		fmt.Print("After 3\n")
	}()

	time.Sleep(5 * time.Second)
	x := <-c

	fmt.Println("Received a signal after 10 sec", x)
}
