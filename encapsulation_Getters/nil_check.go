package main

import "fmt"

func main() {
	// p := &Person{age: 25, name: "Nikhil"}
	var p *Person
	fmt.Println(p.getAge())
	fmt.Println(p.age) // will give panic nil-pointer
}

type Person struct {
	age  int
	name string
}

func (p *Person) getAge() int {
	if p != nil {
		return p.age
	}
	return 0
}

// same as above except return type
// func (p *Person) getAge() (i int) {
// 	if p != nil {
// 		i = p.age
// 	}
// 	return
// }
