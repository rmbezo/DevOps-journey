package main

import "fmt"

func main() {
	// 🔋
	// 🐤
	score := 0

	fmt.Println("Get ready")
	fmt.Println("Score: ", score)
	fmt.Println("")
	for i := 0; i <= 10; i++ {
		fmt.Println("-------------")
		fmt.Println("You are flying to the pipe!", i)
		fmt.Println("🐤 🔋🔋")
		fmt.Println("")

		fmt.Println("You are flying over the pipe!", i)
		fmt.Println(" 🔋 🐤 🔋")
		fmt.Println("")

		fmt.Println("You are flight over the pipe!", i)
		fmt.Println("🔋🔋 🐤")
		fmt.Println("")

		score++
		fmt.Println("Score: ", score)
		fmt.Println("")
	}
}
