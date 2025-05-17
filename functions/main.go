package main

import "fmt"

func main() {
	foo()
	// 1 parameter no return
	bar("tharika")
	// 1 parameter one return
	s := loca("tharika")
	fmt.Println(s)
	// passing 2 parameter and 2 returns
	s1, n := dogYears("ammu", 3)
	fmt.Println(s1, n)
}

func foo() {
	fmt.Println("Iam from foo")
}
func bar(s string) {
	fmt.Println("my name is: ", s)
}

// using sprint because of s
func loca(s string) string {
	return fmt.Sprint("They call me:", s)
}
func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}

//passing no parameter and no return in function
//passing 1 parameter and no return
// passing 2 parameter and 2 returns



// func(r receiver)(identifier(p parameters)) (return(S)) {   code  }