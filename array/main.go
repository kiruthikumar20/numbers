package main

import "fmt"

func main() {
	a := [5]int{}
	for i := 0; i < 5; i++ {
		a[i] = i
	}
	for i, v := range a {
		fmt.Println(i, v)
	}

}
