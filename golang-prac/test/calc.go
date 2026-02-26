package main 

import "fmt"

func main() {
	var pause string
	var num1, num2 float64
	var operator string
	fmt.Println("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Println("Введите оператор(+ - : *): ")
	fmt.Scan(&operator)
	fmt.Println("Введите второе число: ")
	fmt.Scan(&num2)

	// calculation
	var result float64
	switch operator {
	case "+":
			result = num1 + num2 
	case "-":
			result = num1 - num2 
	case ":":
			if num2 == 0 {
				fmt.Println("Делить на ноль нельзя!")
				fmt.Println("Введите ENTER для завершения программы: ")
				fmt.Scanln(&pause)
				return
			} else {
					result = num1 / num2 
			}
	case "*":
			result = num1 * num2 
	default:
			fmt.Println("Non an operator")
			fmt.Println("Введите ENTER для завершения программы: ")
			fmt.Scanln(&pause)
			return
	}

	// The result
	fmt.Println(num1, operator, num2, "=", result)
	fmt.Println("Введите ENTER что бы завершить программу: ")
	fmt.Scanln(&pause)
}
