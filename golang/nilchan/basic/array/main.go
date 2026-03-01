package main

import "fmt"

func main() {
	var hw [2]string
	hw[0] = "Hello"
	hw[1] = "World"
	fmt.Println(hw)

	primes := [6]int{1, 2, 4, 6, 7}
	fmt.Println(primes)
	fmt.Println(primes[3])
}
