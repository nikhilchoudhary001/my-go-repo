package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	//fmt.Println(time.Now().Before(time.Now().Sub()(1 * time.Minute)))
	//p := new(chan int)
	c := make(chan string)
	for i := 0; i < 5; i++ {
		go func(i int) {
			<-c
			c <- fmt.Sprintf("goroutine %d", i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	c <- "hi"
	fmt.Println(<-c)

	// tm := time.Now()
	// if tm.Before(tm.Add(50 * time.Second)) {
	// 	fmt.Println("test")
	// } else {
	// 	fmt.Println("time taken more than 2 sec")
	// }

	x := 5
	fmt.Println(int64(time.Duration(x) * time.Second))

	fmt.Println(uint8(255))
	fmt.Println(math.Pow(2, 15))

}
