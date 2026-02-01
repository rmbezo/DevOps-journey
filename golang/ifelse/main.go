package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("Hi kid")
	} else if age < 50 {
		fmt.Println("Yo wassup bro!")
	} else {
		fmt.Println("Respect for you!")
	}
}
