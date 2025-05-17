package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My dog is walking:", d.first)
}
func (d *dog) run() {
	d.first = "Hennry"
	fmt.Println("My dog is running:", d.first)
}

type dogIn interface {
	walk()
	run()
}

func dogOut(s dogIn) {
	s.run()
	s.walk()
}

func main() {
	d1 := dog{"Ammu"}
	d1.walk()
	d1.run()
	//	dogOut(d1.first)

	d2 := &dog{"Rose"}
	d2.walk()
	d2.run()
	dogOut(d2)
}
