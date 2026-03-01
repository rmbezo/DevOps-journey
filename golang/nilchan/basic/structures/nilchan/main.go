package main

import "fmt"

type User struct {
	Name   string
	Rating float64
	Visits int
}

func main() {
	user := User{
		Name:   "Petya",
		Rating: 33.2,
		Visits: 1,
	}

	fmt.Println("Visits before:", user.Visits)
	Greeting(&user)
	fmt.Println("Visits after:", user.Visits)
}

func Greeting(u *User) {
	fmt.Println("Hello everybody!")
	fmt.Println("My name is:", u.Name)
	fmt.Println("My rating is:", u.Rating)

	u.Visits += 1
}
