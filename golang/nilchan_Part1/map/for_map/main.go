package main

import "fmt"

func main() {
	weather := map[int]int{
		10: 12,
		11: 13,
		12: 11,
		13: 100,
	}

	fmt.Println("Before: ")
	fmt.Println(weather)
	for k, _ := range weather {
		weather[k] += 1
	}
	fmt.Println("After: ")
	fmt.Println(weather)
}
