package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(a...)
	fmt.Println("the sum is:", x)
}

//using veridatic function

func sum(i ...int) int {
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	n := 0
	for _, v := range i {
		n += v
	}
	return n
}
