package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	sliceInt := []int{}
	sliceString := []string{"f", "d", "b", "e", "z", "j", "x", "p", "l", "q", "a", "s"}

	randInt := rand.IntN(100)
	for i := 0; i < randInt; i++ {
		randPlus := rand.IntN(80)
		sliceInt = append(sliceInt, i+randPlus)
	}

	// Before sort
	fmt.Println("Before sort: ")
	fmt.Println(sliceString)
	fmt.Println(sliceInt)

	bubble_sortInt(sliceInt)
	bubble_sortString(sliceString)

	// After start
	fmt.Println("After sort: ")
	fmt.Println(sliceString)
	fmt.Println(sliceInt)

}

// O(n^2)
func bubble_sortInt(l []int) {
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l)-1-i; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
}

// O(n^2)
func bubble_sortString(l []string) {
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l)-1-i; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
}
