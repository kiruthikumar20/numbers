package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("My book name:", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("this is page number:", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("Log From 456", s.String())
}
func main() {
	b := book{
		title: "The Dark Knight",
	}
	var n count = 23

	logInfo(b)
	logInfo(n)
}
