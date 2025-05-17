package main

import (
	"fmt"
	"sort"
)

// type Person struct {
// 	First string
// 	Age   int
// }
// type ByName []Person

// func (a ByName) Len() int           { return len(a) }
// func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

// func main() {
// 	p1 := Person{"kiruthika", 29}
// 	p2 := Person{"tharika", 4}
// 	p3 := Person{"shalini", 24}
// 	p4 := Person{"nadhi", 27}

// 	people := []Person{p1, p2, p3, p4}
// 	fmt.Println(people)
// 	sort.Sort(ByName(people))
// 	fmt.Println(people)

// a := []int{5, 6, 3, 6, 8, 3, 9, 0, 3, 6, 7, 5, 1}
// b := []string{"ad", "gf", "de", "ger", "rtewgre", "bsdfsdf"}
// fmt.Println(a)
// sort.Ints(a)
// fmt.Println(a)
// fmt.Println(b)
// sort.Strings(b)
// fmt.Println(b)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)
	
}
