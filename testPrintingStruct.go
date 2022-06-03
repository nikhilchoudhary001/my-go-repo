package main

import "fmt"

type Customer1 struct {
	name string
	age  int
}

func (c *Customer1) String() string {
	return fmt.Sprintf("Name is %s and age is %d", c.name, c.age)
}

func main() {
	c := new(Customer1)
	c.name = "Nikhil"
	c.age = 28
	fmt.Printf("Customer1 %s", c)
}
