package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func add(a int, b int) int {
	return a + b
}
func addI(a float64, b float64) float64 {
	return a + b
}

// type set interface
type Numbers interface {
	constraints.Integer | constraints.Float
}

// this is type constraint adding type to the function
func addF[T Numbers](a, b T) T {
	return a + b
}

//type alias creating a type and assigning it to a variable

type MyNumbers int

func main() {
	var a MyNumbers = 35
	d := add(4, 5)
	fmt.Println(d)
	g := addI(3.5, 6.7)
	fmt.Println(g)
	fmt.Println(addF(a, 4))
	fmt.Println(addF(2.5, 4.5))

}
