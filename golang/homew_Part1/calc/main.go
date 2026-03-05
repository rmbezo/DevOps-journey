package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(" ")
		fmt.Println("Hello, im your go calculator")

		fmt.Print("Type the problem (12 + 2, 14 / 2, 99 - 9): ")
		scanner.Scan()

		input := strings.Fields(scanner.Text())
		if len(input) != 3 {
			fmt.Println("Ошибка: введите выражение целиком (число знак число)")
			continue
		}
		num1, err := strconv.ParseFloat(input[0], 64)
		if err != nil {
			fmt.Println("First number is invalid:", input[0])
			continue
		}
		num2, err := strconv.ParseFloat(input[2], 64)
		if err != nil {
			fmt.Println("Second number is invalid:", input[2])
			continue
		}
		var result float64

		//Count
		switch input[1] {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "/", ":":
			if num2 == 0 {
				fmt.Println(" ")
				fmt.Println("Impossible to divide on zero!")
				fmt.Println(" ")
				continue
			} else {
				result = num1 / num2
			}
		case "*":
			result = num1 * num2
		case "**", "^":
			result = math.Pow(num1, num2)
		default:
			fmt.Println("Not a operator!")
			continue
		}
		fmt.Println(" ")
		fmt.Printf("The result is: %.5f %v %.5f = %.5f", num1, input[1], num2, result)
		fmt.Println(" ")
	}

}
