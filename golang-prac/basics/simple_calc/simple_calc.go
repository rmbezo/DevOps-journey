package main

import "fmt"

func main() {
	var num1, num2 float64 // Используем float64 для дробных чисел
	var operator string

	fmt.Println("--- Простой Калькулятор на Go ---")

	// 1. Ввод первого числа
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&num1)

	// 2. Ввод оператора
	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)

	// 3. Ввод второго числа
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&num2)

	// 4. Выполнение операции
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 { // Проверка деления на ноль
			result = num1 / num2
		} else {
			fmt.Println("Ошибка: Деление на ноль невозможно!")
			return // Выходим, если ошибка
		}
	default:
		fmt.Println("Ошибка: Неизвестный оператор!")
		return // Выходим, если ошибка
	}

	// 5. Вывод результата
	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
