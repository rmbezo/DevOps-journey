package main

import "fmt"

func main() {
	number := 15

	var ptr *int = &number

	fmt.Println("number:", number)
	fmt.Println("Pointer:", ptr)

	if ptr != nil {
		fmt.Println("Unnamed:", *ptr)
	} else {
		fmt.Println("Nil pointer.")
	}
	// fmt.Println("Unnamed:", *ptr) ## ERROR
}
