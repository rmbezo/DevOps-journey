package main

import "fmt"

func main() {
	name := "Gosha"
	ptr := &name

	fmt.Println("Before changing: ", name)
	change_name(ptr)
	fmt.Println("Name after changing: ", name)
}

func change_name(n *string) {
	*n = "Lesha"
}
