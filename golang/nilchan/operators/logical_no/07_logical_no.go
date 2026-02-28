package main

import "fmt"

func main() {
	subscribe := true

	if !subscribe {
		fmt.Println("You are not subscribed")
	} else {
		fmt.Println("You are already subscribed")
	}
}
