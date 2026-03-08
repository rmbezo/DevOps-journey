package main

import "fmt"

type Array struct {
	num int
}

func main() {
	array := [10]int{5, 3, 2, 9, 8, 2, 1, 15, 12, 14}

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}
