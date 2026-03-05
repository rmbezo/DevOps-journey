package main

import "fmt"

func main() {
	number := 10
	pointer := &number
	test(pointer)
}

func test(x *int) {
	fmt.Println(x)
}
