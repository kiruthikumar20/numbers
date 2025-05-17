package main

import "fmt"

func main() {
	a := []int{23, 24, 25, 26}
	fmt.Println(a)
	av := append(a, 27, 28)
	fmt.Println(av[:2])
}
