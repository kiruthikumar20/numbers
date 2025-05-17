package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Age   int
}

func main() {
	p1 := Person{
		First: "tharika",
		Age:   4,
	}
	p2 := Person{
		First: "henry",
		Age:   24,
	}
	p3 := Person{
		First: "James",
		Age:   54,
	}
	people := []Person{p1, p2, p3}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(bs))
}
