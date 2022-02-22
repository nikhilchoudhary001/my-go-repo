package main

import (
	"fmt"
)

func main() {
	i := 0
	// panic: too many concurrent operations on a single file or socket (max 1048575)
	for {
		i++
		go test("Nikhil" + fmt.Sprint(i))
	}

	// Below line will not be reached
	//time.Sleep(10 * time.Second)
}

func test(s string) {
	fmt.Println("s is ", s)
}
