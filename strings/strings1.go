package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcd!ef"

	// for i, c := range s {
	// 	fmt.Printf("Char - %d is %c \n", i, c)
	// }

	fmt.Println(string(s[0]))
	fmt.Println(s[0] - 'a')
	fmt.Println(s[1:4])
	fmt.Println(strings.Split(s, "!"))

	// cannot assign a string at a particular index(datatype of byte)// will give compile time err as we are assigning
	// s[0] = "A"

}
