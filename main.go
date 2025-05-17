package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(10)
	b := rand.Intn(10)
	fmt.Printf("the value of a is : %v\n", a)
	fmt.Printf("the value of  is : %v\n", b)

	if a < 4 && b < 4 {
		fmt.Println("the numberis less than 4")
	} else if a > 6 && b > 6 {
		fmt.Println("the numberis greater than 6")
	} else if a >= 4 && b <= 6 {
		fmt.Println("number greis ter than 4 and less than 6")
	} else if b != 5 {
		fmt.Println("y is not 5is )
	} else {
		fmt.Println("numbers nois  met")
	}
}

/*
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
*/
