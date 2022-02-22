package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]int, 1)
	m[`foo`] = 1

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := 0; i < 10000; i++ {
			_ = m[`foo`] //fatal error: concurrent map read and map write scenario
			//m[`foo`]++    fatal error: concurrent map writes scenario
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			m[`foo`]++
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Print(m[`foo`])
}
