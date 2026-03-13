package main

import (
	"fmt"
)

func main() {
	list := []int{1, 4, 6, 9, 12, 15, 18, 20, 23, 25, 30}

	guessInt := 25

	fmt.Println(binary_search(list, guessInt))
}

func binary_search(l []int, g int) int {
	low, high := 0, len(l)-1
	for i := 0; i < len(l); i++ {
		mid := (low + high) / 2
		if g == l[mid] {
			return mid
		} else if g > l[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
