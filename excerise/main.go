package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	delete(m, "fleming_ian")

	fmt.Println("the m contains:m", m)

	fmt.Println("---------------------------------")
	for k, v := range m {
		fmt.Println(k, v)
		fmt.Println("---------------------------------")
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
