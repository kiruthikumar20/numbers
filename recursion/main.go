package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(4))
	fmt.Println(factLoop(4))

}
func factorial(n int) int {
	fmt.Println("the number of n:", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
		fmt.Println("the number of n:", n)
	}
	return total
}

/*
the function calls itself is called recursion. Here we are creating a func factorial
passing a parameter and return should be in int, then we are telling if n == 0 i need
result 1 or run n * function name minus 1
*/
