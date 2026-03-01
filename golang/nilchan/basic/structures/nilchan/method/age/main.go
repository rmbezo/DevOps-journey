package main

import "fmt"

type User struct {
	age int
}

func main() {
	user := &User{
		age: 12,
	}
	fmt.Println("User in begining:", user.age)
	user.changeage1(33)
	fmt.Println("User after change 1:", user.age)
	user.changeage2(24)
	fmt.Println("User after change 2:", user.age)
}

func (u User) changeage1(newage int) {
	u.age = newage
}
func (u *User) changeage2(newage int) {
	u.age = newage
}
