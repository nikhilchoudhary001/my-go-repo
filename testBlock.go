package main

import (
	"fmt"
	"time"
)

func funtest() {
	fmt.Println("f")
}

func main() {
	{ // start outer block
		a := 1
		fmt.Println(a)
		{ // start inner block
			b := 2
			fmt.Println(a)
			fmt.Println(b)
		} // end inner block
	} // end outer block

	// for i := 0; i < 10; i++ {
	// 	//fmt.Printf("%p \n", c)
	// 	fmt.Print("N", i)
	// 	i := i
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()
	// }

	if i := 5; i >= 0 {
		fmt.Println(i)
		var i int = 7
		fmt.Println(i)

	}

	switch i := 2; i * 4 {
	case 8:
		j := i * 2
		fmt.Println(j)
	default:
		fmt.Println("default")
	}

	//time.Sleep(5 * time.Second)

	tick := time.Tick(100 * time.Millisecond)
LOOP:
	for {
		select {
		case <-tick:
			i := 0
			fmt.Println("tick", i)
			break LOOP
		default:
			// "i" is undefined here
			time.Sleep(30 * time.Millisecond)
			fmt.Println("sleep")
		}
	}
}
