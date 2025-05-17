package main

import "fmt"

func intDelta(n *int) {
	*n = 45
}
func sliceData(ii []int) {
	ii[2] = 30
}
func mapData(s map[string]int, t string) {
	s[t] = 5
}
func main() {
	a := 10
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)
	xi := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(xi)
	sliceData(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["tharika"] = 4
	fmt.Println(m["tharika"])
	mapData(m, "tharika")
	fmt.Println(m["tharika"])

	// a := 67
	// fmt.Println(a)
	// intDelta(&a)
	// fmt.Println(a)

	// a := 10
	// fmt.Println(a)
	// fmt.Printf("the value of a : %v and the type is %T\n", a, &a)
	// y := &a
	// fmt.Println(*y)
	// fmt.Println(y)
	// *y = 43
	// fmt.Println(a)
}

/*
here we are assigning a "a" variable and giving value as 10, then printing it to
check type and value
again variable y assiging address to value a it will go the a variable and
take that value
in println we are using pointers ----> pointer will go the address of memory and that
take that value and it will print

*/
