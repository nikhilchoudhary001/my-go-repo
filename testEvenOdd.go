package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 1000

	ch := make(chan bool)
	num := make(chan string)
	done := make(chan bool)

	go func() {
		for i := 1; i < number; i = i + 2 {
			if i == 1 {
				num <- "Odd " + strconv.Itoa(i)
				ch <- true
			} else {
				<-ch
				num <- "Odd " + strconv.Itoa(i)
				ch <- true
			}
		}
	}()

	go func() {
		for i := 2; i <= number; i = i + 2 {
			if i == number {
				<-ch
				num <- "Even " + strconv.Itoa(i)
				close(num)
				close(ch)
			} else {
				<-ch
				num <- "Even " + strconv.Itoa(i)
				ch <- true
			}

		}
		done <- true
	}()

	go func() {
		for i := range num {
			fmt.Println(i)
		}
	}()

	<-done
	fmt.Println("Finished")
}
