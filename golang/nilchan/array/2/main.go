package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	userArray := [3]User{
		User{
			Name:    "Gosha",
			Rating:  12.5,
			Premium: false,
		},
		User{
			Name:    "Victoria",
			Rating:  59.5,
			Premium: true,
		},
		User{
			Name:    "Mary",
			Rating:  9.5,
			Premium: false,
		},
	}
	// Pointer on array for functions
	userArray_ptr := &userArray

	fmt.Println("Before changes:")
	for i := 0; i < len(userArray); i++ {
		pp.Println(userArray[i])
	}

	Premium_Rating(userArray_ptr)

	fmt.Println("")
	fmt.Println("After changes: ")
	for i := 0; i < len(userArray); i++ {
		pp.Println(userArray[i])
	}
}

func Premium_Rating(u *[3]User) {
	for i := 0; i < len(u); i++ {
		if u[i].Premium {
			u[i].Rating += 1
		}
	}
}
