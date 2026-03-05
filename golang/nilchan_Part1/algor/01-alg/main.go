package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Privet Vasiliy!")
	var b int = rand.IntN(1000)
	var a int = rand.IntN(b)
	fmt.Printf("a: %d ; b: %d\n", a, b)
	for b > a {
		if b%10 == 1 {
			b = (b - 1) / 10
			fmt.Println("cut 1 ->", b)
		} else if b%2 == 0 {
			b = b / 2
			fmt.Println("/2 ->", b)
		} else {
			fmt.Println("Impossible!")
			return
		}
	}

	if b == a {
		fmt.Println("YES! We reached a")
	} else {
		fmt.Println("NO :(")
	}
}
