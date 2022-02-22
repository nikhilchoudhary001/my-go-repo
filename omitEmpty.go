package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Age *string `json:",omitempty"`
}

func main() {
	age := ""

	d := Dog{Age: &age}

	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}
