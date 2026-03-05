package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name   string
	Age    int
	Rating float64
}

func main() {
	fmt.Println("Hello, World!")

	user := []User{
		User{
			Name:   "Gosha",
			Age:    12,
			Rating: 55.0,
		},
		User{
			Name:   "Roma",
			Age:    17,
			Rating: 923.0,
		},
		User{
			Name:   "Vova",
			Age:    52,
			Rating: 52.0,
		},
	}
	//Pointer
	userArray := &user

	for index, user := range *userArray {
		pp.Println(index, user)
	}
}
