package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49}
	fmt.Println(x)

	x = append(x, 50)
	fmt.Println(x)

	x = append(x, 51, 52)
	fmt.Println(x)
}
