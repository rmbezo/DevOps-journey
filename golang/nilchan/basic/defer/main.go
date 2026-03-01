package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()
	defer func() {
		fmt.Println("defer 3")
	}()
	defer func() {
		fmt.Println("defer 4")
	}()

	fmt.Println("Before defer_privet")
	defer_privet()
	fmt.Println("After defer privet")
}

func defer_privet() {
	for i := 0; i <= 4; i++ {
		defer func() {
			fmt.Println("defer ", -i)
		}()
	}
}
