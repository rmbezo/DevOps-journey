package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hello, im your go calculator")

	fmt.Print("Type your first number: ")
	scanner.Scan()
	input1 := strings.TrimSpace(scanner.Text())
	num1, _ := strconv.ParseFloat(input1, 64)

	fmt.Print("Type your second number: ")
	scanner.Scan()
	input2 := strings.TrimSpace(scanner.Text())
	num2, _ := strconv.ParseFloat(input2, 64)

	res1 := num1 + num2
	res2 := num1 - num2

	fmt.Printf("Sum: %.2f, Minus: %.2f", res1, res2)
}
