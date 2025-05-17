package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	logIn := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(logIn))
	if err != nil {
		fmt.Println("password is incorrect")
	}
	fmt.Println("logged in")
}
