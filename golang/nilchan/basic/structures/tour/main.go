package main

import "fmt"

type Account struct {
	Subscribers int
	Premium     bool
	Subscribed  int
}

func main() {

	// Account
	user := Account{1241, true, 9}
	fmt.Printf("You have %v subscribers\n", user.Subscribers)
	if user.Premium {
		fmt.Println("You have Premium subscription")
	} else {
		fmt.Println("You dont have Premium subscription")
	}
	fmt.Printf("You subscribed to %v channels", user.Subscribed)
}
