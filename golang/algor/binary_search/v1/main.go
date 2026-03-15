package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	scanner := buffio.NewScanner(os.Stdin)
}
