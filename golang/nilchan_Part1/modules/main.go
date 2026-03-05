package main

import (
	"fmt"
	"mod-test/poka"
	"mod-test/privet"
)

func main() {
	fmt.Println("Package privet:")
	privet.Privet()

	fmt.Println("Package privet:")
	poka.Poka(50)

}
