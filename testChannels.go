package main

import (
	"fmt"
	"time"
)

func main() {

	quit := make(chan bool)
	go func() {
		for {
			select {
			case i := <-quit:
				fmt.Println("Received quit signal", i)
				return
			default:
				fmt.Println("Hello world")
			}
		}
	}()

	fmt.Println("Program Started")
	time.Sleep(1 * time.Millisecond)
	quit <- true
	time.Sleep(1 * time.Millisecond)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
