package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Type your name: ")
	scanner := bufio.NewScanner(os.Stdin)

	if ok := scanner.Scan(); !ok {
		fmt.Println("Error")
		return
	}

	Text := scanner.Text()

	fmt.Println("Your name is: ", Text)
}
