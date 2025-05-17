package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m := map[string]int{
		"james": 42,
		"Money": 32,
	}
	for key, v := range m {
		fmt.Println("range of value: m", key, v)
	}
	age1 := m["james"]
	fmt.Println("the age of james:", age1)

	if v, ok := m["james"]; ok {
		fmt.Println(v)
	}
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 4 {
			fmt.Println("x is 3")
		}
	}
}
