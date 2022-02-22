package main

import (
	"fmt"
	"io/ioutil"
)

// https://gobyexample.com/reading-files
func main() {

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("File contents - ", string(data))
}
