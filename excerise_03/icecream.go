package main

import "fmt"

type fav struct {
	first    string
	last     string
	iceCream []string
}

func main() {

	p1 := fav{
		first:    "tharika",
		last:     "sudhakar",
		iceCream: []string{"chocolate", "vanilla", "mango"},
	}

	for i, v := range p1.iceCream {
		fmt.Println("my favorite ice cream is:", i, v)
	}

	m := map[string]fav{
		p1.last: p1,
	}
	for _, v := range m {
		fmt.Println(v)
		for _, v := range v.last {
			fmt.Println(p1.last, p1.last, v)
		}
	}
}
