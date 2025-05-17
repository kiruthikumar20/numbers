package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)
}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

/*
func main() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := incrementor()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
*/
