package main

import "fmt"

type Employee struct {
	Name string
	ID   int32
}

func main() {
	//var i interface{}

	i := Employee{Name: "Nikhil Choudhary", ID: 90}

	fmt.Println("Employee = ", fmt.Sprintf("%+v", i))
}
