package main

import (
	"encoding/json"
	"fmt"
)

// type Dog struct {
// 	Age *string `json:",omitempty"`
// }

func main() {
	// age := ""

	// d := Dog{Age: &age}

	// b, _ := json.Marshal(d)
	// fmt.Println(string(b))

	d := Dog{
		Breed: "pug",
	}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}

type dimension struct {
	Height int
	Width  int
}

type Dog struct {
	Breed    string
	WeightKg int        `json:",omitempty"`
	Size     *dimension `json:",omitempty"`
}
