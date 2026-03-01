package main

import "fmt"

func main() {
	sunny := true
	weekend := true

	if sunny && weekend {
		fmt.Println("Let's go for a walk!")
	} else {
		fmt.Println("Not today my friend.")
	}
}
