package main

import "fmt"

func main() {
	score := 25

	if score > 20 {
		fmt.Println("Ti mega krasavchik!!")
	} else if score > 10 {
		fmt.Println("Ti krasava!")
	} else {
		fmt.Println("You have to practice my friend...")
	}
}
