package main

import "fmt"

func main() {
	num := 10
	point := &num

	test(point)

	change_env(point)
	fmt.Println(num)
	fmt.Println("Unname pointer: ", *point)
}

func test(x *int) {
	fmt.Println(x)
	fmt.Println(*x)
}

func change_env(x *int) {
	*x = 500
}
