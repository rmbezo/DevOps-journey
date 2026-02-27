package main

import "fmt"

func main() {
	fmt.Println("Начало функции")
	count := 11
	for i := 0; i < count; i++ {
		fmt.Println("Я знаю цифру:", i)
	}
}
