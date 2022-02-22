package main

import (
	"fmt"
)

func main() {

	// for i := 0; i < 5000; i++ {
	// 	go func() {
	// 		fmt.Printf("Value is %d", i)
	// 	}()
	// }
	// time.Sleep(time.Second)

	messages := make(chan string)

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

}
