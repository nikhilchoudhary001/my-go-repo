package main

// http://www.golangpatterns.info/object-oriented/operators
type Matrix interface{}

type denseMatrix struct{}

type sparseMatrix struct{}

func (a *denseMatrix) Plus(b Matrix) Matrix

func (a *sparseMatrix) Plus(b Matrix) Matrix

func (a *denseMatrix) test() {}
func test()                  {}

func main() {
}
