package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	// Enter the first number
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)

	// Enter the operator
	fmt.Println("Enter the operator(*, :, +, -): ")
	fmt.Scanln(&operator)

	// Enter the second number
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)

	// Calculation part
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Ошибка: Деление на ноль невозможно!")
			return
		}
	default:
		fmt.Println("Ошибка: Неизвестный оператор!")
		return
	}

	// Result
	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
