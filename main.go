package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("the value of x : %v\n", x)
	if x <= 100 {
		fmt.Println("less than 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("greater than 100 and less than 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("greater than 200 and less than 250")
	} else {
		fmt.Println("greater than 300")
	}
}
