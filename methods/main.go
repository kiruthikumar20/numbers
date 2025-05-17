package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("Iam ", p.first)
}
func main() {
	p1 := person{
		first: "tharika",
	}
	p2 := person{
		first: "kiruthika",
	}
	p1.speak()
	p2.speak()
}
