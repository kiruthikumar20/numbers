package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 3))
}

// printSquare is the callback function here
func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number is:%v, and the squared number is:%v", a, x)

}

func square(n int) int {
	return n * n
}
