package main

import "fmt"

type Account struct {
	Subscribers int
	Premium     bool
	Subscribed  int
}

func GivePremium(u *Account) {
	u.Premium = true
	u.Subscribers += 100
}

func main() {
	user := Account{1241, false, 9}

	fmt.Println("Before bonus, premium:", user.Premium)

	GivePremium(&user)

	fmt.Println("After bonus premium:", user.Premium)
	fmt.Printf("Now we have %v subscribers\n", user.Subscribers)
}
