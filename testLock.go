package main

import (
	"fmt"
	"sync"
	"time"
)

var l sync.Mutex
var l2 sync.Mutex

func main() {
	//m := make(map[string]int, 1)
	msg := make(chan int, 5)

	var value int = 0
	// for _, i := range []int{1, 2, 3} {
	// 	fmt.Println(i)
	// }
	//	os.Exit(0)
	go func() {
		for i := 0; i < 5; i++ {
			value = i
			time.Sleep(1 * time.Second)
		}
	}()

	//time.Sleep(1 * time.Second)
	//value = 5
	time.Sleep(1 * time.Second)
	count := 0

	for {
		select {
		case msg <- value:
			fmt.Println("Value consumed")
			count++
		default:
			fmt.Println("Default accoured")
			break
		}
		break
	}

	fmt.Println(count)

}
