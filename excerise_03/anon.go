package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "tharika",
		friends: map[string]int{
			"shalu":   24,
			"nadhi":   27,
			"tharrun": 18,
		},
		favDrinks: []string{"coke", "7up", "pespi"},
	}
	for k, v := range p1.favDrinks {
		fmt.Println(p1.first, "drinks", p1.favDrinks, k, v)
	}
	for k, v := range p1.friends {
		fmt.Println(p1.first, "friends", p1.friends, k, v)
	}
}
