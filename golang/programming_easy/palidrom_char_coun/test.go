package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 522
	s := strconv.Itoa(x)

	fmt.Println("s:", s)
	fmt.Println("s char 1:", s[0])
}
