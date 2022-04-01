package main

import (
	"fmt"
	"sync"
)

type car struct {
	wheel
}

type wheel struct {
}

func (w wheel) airvalue() int {
	return 0
}

type bike struct {
}

type movement interface {
	Move()
}

func (c car) Move() {
	c.airvalue()
}
func (c bike) Move() {

}

var n int = 10
var wg sync.WaitGroup

func main() {
	var integers = []int{1, 2, 3, 4, 5}
	fmt.Println(integers)
	// temp:= integers
	integers = integers[:1]
	// integers=append(integers, temp[2:])

	chann := make(chan int)
	wg.Add(1)
	go printOdd(chann)
	go printEven(chann)
	wg.Wait()
}

func moved(m movement) {

}

func printOdd(c chan int) {
	i := 1
	for {
		if i != 1 {
			i = <-c
		}
		if i > 10 {
			wg.Done()
			break
		}
		fmt.Printf("Go routine: odd :%d\n", i)
		i++
		c <- i
	}
}

func printEven(c chan int) {
	for {
		i := <-c
		if i > 10 {
			wg.Done()
			break
		}
		fmt.Printf("Go routine: even :%d\n", i)
		i++
		c <- i
	}
}
