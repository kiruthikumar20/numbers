package main

import "fmt"

type engine struct {
	electric bool
}

type vechile struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	v1 := vechile{
		engine: engine{
			electric: true,
		},

		make:  "bmw",
		model: "Mustang",
		doors: 4,
		color: "blue",
	}
	v2 := vechile{
		engine: engine{
			electric: false,
		},
		make:  "audi",
		model: "Mukad",
		doors: 3,
		color: "white",
	}
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.electric, v1.make, v1.model)
	fmt.Println(v2.electric, v2.make, v2.color)

}
