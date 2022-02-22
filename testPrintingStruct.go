package main

import "fmt"

type Customer struct {
	name string
	age  int
}

func (c *Customer) String() string {
	return fmt.Sprintf("Name is %s and age is %d", c.name, c.age)
}

func main() {
	c := new(Customer)
	c.name = "Nikhil"
	c.age = 28
	fmt.Printf("Customer %s", c)
}
