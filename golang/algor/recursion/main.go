package main

import "fmt"

func main() {
	fmt.Println(factorial(3))
}

func factorial(x int) int {
	fmt.Println(x)
	if x == 1 {
		return x
	} else {
		x = x * factorial(x-1)
		return x
	}
}
