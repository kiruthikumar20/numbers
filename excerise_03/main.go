package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "tharika",
		last:  "sudhakar",
		age:   4,
	}

	p2 := person{
		first: "kiruthika",
		last:  "kumar",
		age:   29,
	}
	fmt.Printf("name and age: %v\n", p1)
	fmt.Printf("name and age: %v\n", p2)

	fmt.Println(p1.first, p1.last, p1.age)
	

}
