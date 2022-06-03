package main

import "fmt"

func main() {
	//p := &Person{age: 25, name: "Nikhil"}
	var p *Person
	fmt.Println(p.getAge())
	fmt.Println(p.age)
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
