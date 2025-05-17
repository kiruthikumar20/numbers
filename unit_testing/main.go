package main

import "fmt"

func main() {
	x := doMath(43, 45, Add)
	fmt.Println(x)
	y := doMath(69, 95, Sub)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func Add(a int, b int) int {
	return a + b
}
func Sub(a int, b int) int {
	return a - b
}
