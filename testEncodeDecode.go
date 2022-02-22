package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
)

func main() {

	num := rand.Uint64()
	// fmt.Println("New number to be encoded ", num)
	// fmt.Println("Length of alphabet ", length)
	// fmt.Println("Mod value ", 5577006791947779410%length)
	// fmt.Println("Index value ", alphabet[40])
	// fmt.Println("Divide value ", 5577006791947779410/length)

	c := 'O' // rune (characters in Go are represented using `rune` data type)
	asciiValue := int(c)

	fmt.Printf("Ascii Value of %c = %d\n", c, asciiValue)
	str := Encode(num)

	val, _ := Decode(str)
	fmt.Printf("Decoded number %c = %d\n", str, val)
}

func Encode(number uint64) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(11)

	for ; number > 0; number = number / length {
		fmt.Println("Inside loop ", number)
		encodedBuilder.WriteByte(alphabet[(number % length)])
	}

	fmt.Println("After encoding ", encodedBuilder.String())
	return encodedBuilder.String()
}

func Decode(encoded string) (uint64, error) {
	var number uint64

	for i, symbol := range encoded {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)

		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}
