package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)

	fmt.Print("Введите оператор(+ - : *): ")
	fmt.Scan(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	var result float64
	switch operator {
	case "+":
		{
			result = num1 + num2
		}
	case "-":
		{
			result = num1 - num2
		}
	case ":":
		{
			if num2 == 0 {
				fmt.Println("Делить на ноль нельзя!")
				return
			} else {
				result = num1 / num2
			}
		}
	case "*":
		{
			result = num1 * num2
		}
	default:
		fmt.Println("Неизвестный оператор!")
		return
	}
	fmt.Println(num1, operator, num2, "=", result)

	if result > 100 {
		more()
	} else {
		less()
	}
}

func more() {
	fmt.Println("Ваше число больше 100")
}

func less() {
	fmt.Println("Ваше число меньше 100")
}
