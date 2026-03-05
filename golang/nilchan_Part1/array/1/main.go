package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Before for:", arr)
	for i := 0; i < 5; i++ {
		if arr[i]%2 == 0 {
			arr[i] *= 2
		}
	}
	fmt.Println("After for:", arr)

	for i := 0; i < 5; i++ {
		fmt.Println(i, "-", arr[i])
	}
}
