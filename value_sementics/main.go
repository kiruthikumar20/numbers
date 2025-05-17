package main

import "fmt"

func addOne(i int) int {
	return i + 1
}
func addOneP(v *int) {
	*v += 1
}

func main() {
	// value semantics
	a := 1
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)
	// pointer semantics
	b := 1
	fmt.Println("value semantics")
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)

}
