package main

import "fmt"

func main() {
	// Name - prisoner
	// int - how many years left
	Prison := make(map[string]float64, 5)
	Prison["Gosha"] = 12.0
	Prison["Jora"] = 0.5
	Prison["Jenya"] = 10.90
	Prison["Pasha"] = 15.90
	Prison["Mark"] = 5.0

	fmt.Println("Before breaking the limits: ")
	fmt.Println(Prison)

	Prison["Yarik"] = 99.0
	fmt.Println("After breaking the limits: ")
	fmt.Println(Prison)
}

// So map have flexible limits
