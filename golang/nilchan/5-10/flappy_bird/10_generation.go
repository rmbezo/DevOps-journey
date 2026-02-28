package main

import "fmt"

func main() {
	fmt.Println("Starting generation of the world...")
	fmt.Println("")

	for i := 1; i <= 5; i++ {
		fmt.Println("-------")
		fmt.Println("Pipe number: ", i)
		fmt.Println("-------")

		if i%2 == 0 {
			fmt.Println("🛢🛢")
		} else {
			fmt.Println("🔋🔋")
		}

		fmt.Println("-------")
		fmt.Println("")
	}
}
