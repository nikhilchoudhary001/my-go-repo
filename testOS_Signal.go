package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	c := make(chan int, 1)

	arr := []int{1, 2}
	// Running the HTTP server
	go func() {
		for _, i := range arr {
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("Number %d \n", i)
		}
		os.Exit(0)
		// time.Sleep(10 * time.Second)
		// c <- 500
	}()

	interruptSignal := <-interrupt
	switch interruptSignal {
	case os.Kill:
		fmt.Println("Got SIGKILL...")
		c <- 100
	case os.Interrupt:
		fmt.Println("Got SIGINT...")
		//time.Sleep(4 * time.Second)
		c <- 400
	case syscall.SIGTERM:
		fmt.Println("Got SIGTERM...")
		c <- 700
	}

	fmt.Println("The service is shutting down...")
	//server.Shutdown(context.Background())
	x := <-c
	fmt.Println("Shut down is done", x)
}
