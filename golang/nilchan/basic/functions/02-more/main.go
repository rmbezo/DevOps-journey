package main

import "fmt"

func main() {
	var n1, n2 float64
	fmt.Print("Type first number: ")
	fmt.Scan(&n1)
	fmt.Print("Type second number: ")
	fmt.Scan(&n2)
	fmt.Println("Summary: ", env(n1, n2))
}

func env(x, y float64) float64 { return x + y }
