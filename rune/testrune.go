package main

import "fmt"

// https://golangbyexample.com/go-index-character-string/
func main() {
	sample := "ab£c"
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%c\n", sample[i])
	}
	fmt.Printf("Length is %d\n", len(sample))
	fmt.Println("Byte sequence is - ", []byte(sample))

	sampleRune := []rune(sample)

	fmt.Printf("%c\n", sampleRune[0])
	fmt.Printf("%c\n", sampleRune[1])
	fmt.Printf("%c\n", sampleRune[2])
	fmt.Printf("%c\n", sampleRune[3])

	for _, c := range sampleRune {
		fmt.Printf("%c\n", c)
	}
}
