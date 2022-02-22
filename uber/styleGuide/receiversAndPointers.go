package main

import "fmt"

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {
	fmt.Println("Inside s1 func")
}

type S2 struct{}

func (s *S2) f() {
	fmt.Println("Inside s2 func")
}

func main() {
	//s1Val := S1{}
	s1Ptr := &S1{}
	//s2Val := S2{}
	//s2Ptr := &S2{}

	var i *S1
	//i = s1Val
	i = s1Ptr
	//i = s2Ptr

	//fmt.Println(s2Val)
	i.f()

	// The following doesn't compile, since s2Val is a value, and there is no value receiver for f.
	//   i = s2Val
}
