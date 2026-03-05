package main

import "fmt"

type User struct {
	Name   string
	Rating float64
}

func main() {
	user := User{
		Name: "Petya",
	}
	user.Greeting()
	user.Goodbye()
}

func (u *User) Greeting() {
	fmt.Println("Hello everybody!")
	fmt.Println("My name is:", u.Name)
	fmt.Println("My rating is:", u.Rating)
	// Someone liked Petya
	u.Rating += 1
}

func (u *User) Goodbye() {
	fmt.Println("Goodbye everybody!")
	fmt.Println("My name is:", u.Name)
	fmt.Println("My rating was:", u.Rating)
}
