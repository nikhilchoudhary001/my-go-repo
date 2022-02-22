package main

import (
	"fmt"
	"time"
)

func main() {
	expiryTime := time.Now().Add(2 * time.Second)
	//fmt.Println("tm before is ", tm)
	time.Sleep(3 * time.Second)
	if expiryTime.Before(time.Now()) {
		fmt.Println("time has expired")
	} else {
		fmt.Println("time has not expired")
	}
}
