package main

import (
	"fmt"
	"time"
)

var sem = make(chan int, 10)

//var c chan *Request = make(chan *Request, 10)

type Request struct {
	a int
}

func handle(r *Request) {
	sem <- 1
	fmt.Println("Inside ", r.a)
	// Wait for active queue to drain.
	//time.Sleep(1 * time.Second) // May take a long time.
	<-sem // Done; enable next request to run.
}

func Serve(queue chan *Request) {
	for req := range queue {
		fmt.Println(req.a)
		go handle(req) // Don't wait for handle to finish.
	}
}

func main() {

	c := make(chan *Request, 100)
	for i := 0; i < 100; i++ {
		req := &Request{i}
		c <- req
	}
	go Serve(c)
	close(c)
	time.Sleep(3 * time.Second)
	fmt.Println("End")

}
