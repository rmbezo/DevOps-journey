package main

import (
	"fmt"
	"math"
)
var num1, num2, result float64
var op string

func plus(x, y float64) float64 { return x + y }
func minus(x, y float64) float64 { return x - y }
func umnojit(x, y float64) float64 { return x * y }
func split(x, y float64) float64 { return x / y }
func stepen(x, y float64) float64 { return math.Pow(x, y) }


func main() {
	fmt.Print("Type the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Type the operator(+, -, : or /, *, ** or /")
	fmt.Scan(&op)
	fmt.Print("Type the second number: ")
	fmt.Scan(&num2)

	switch op {
		case "+":
			result = plus(num1, num2)
		case "-":
			result = minus(num1, num2)
		case ":", "/":
			if num2 != 0 {
				result = split(num1, num2)
			} else {
					fmt.Print("Error! Cannot split on 0! Press ENTER to exit: ")
					fmt.Scanln()
					return
			}
		case "*":
			result = umnojit(num1, num2)
		case "**", "^":
			result = stepen(num1, num2)
		default:
			fmt.Print("Not an operator! Press ENTER to exit: ")
			fmt.Scan()
			return
		}
		fmt.Printf("The result is: %g %s %g = %g\n", num1, op, num2, result)
		fmt.Print("Press ENTER to exit: ")
		fmt.Scan()
}
