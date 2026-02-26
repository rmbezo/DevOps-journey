package main 

import "fmt"

func main() {
	var name string
	var age float64

	fmt.Println("Type your name: ")
	if _, err := fmt.Scan(&name); err != nil {
		fmt.Println("Invalid name!")
		return
	}
	fmt.Println("Type your age: ")
	if _, err := fmt.Scan(&age); err != nil {
		fmt.Println("Invalid age type!")
		return
	}

	fmt.Println("Your name is:", name, ", your age is: ", age)
}
