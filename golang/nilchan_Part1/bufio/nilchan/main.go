package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// Imported
	"github.com/k0kubun/pp"
)

func main() {

	for {
		// Сканим ввод / Scan input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Type command (ENTER - exit): ")

		// Проверяем ок ли он / Checking if its okay
		if ok := scanner.Scan(); !ok {
			fmt.Println("Error")
			return
		}

		// Converting our input into text
		text := scanner.Text()

		// Creating slice
		fields := strings.Fields(text)

		// Checking if len is 0 - exiting program.
		if len(fields) == 0 {
			fmt.Println("Exiting program.")
			return
		}

		// Output the slice + command
		pp.Println(fields)
		fmt.Println("Команда: ", fields[0])

		// Creating second part of outup
		var str string = ""
		// for i := 1; i < len(fields); i++ {
		// 	str += fields[i]

		// 	if i != len(fields)-1 {
		// 		str += " "
		// 	}
		// }
		str = strings.Join(fields[1:], " ")

		// Logical part
		cmd := strings.ToLower(fields[0])
		switch cmd {
		case "добавить":
			fmt.Println("Seems like you want to add:", str)
		case "удалить":
			fmt.Println("Seems like you want to delete:", str)
		case "утилизировать":
			fmt.Println("Seems like you want to burn it:", str)
		case "захавать":
			fmt.Println("Seems like you want to eat:", str)
		case "заказать":
			fmt.Println("Seems like you want to order:", str)
		case "help":
			fmt.Println("Команда: <добавить> {что нужно добавить, напр. хавку}")
			fmt.Println("---- This command let's you add something")
			fmt.Println("")
			fmt.Println("Команда: <удалить> {что нужно удалить}")
			fmt.Println("---- This command let's you delete something")
			fmt.Println("")
			fmt.Println("Команда: <утилизировать> {что нужно утилизировать}")
			fmt.Println("---- This command let's you burn something")
			fmt.Println("")
			fmt.Println("Команда: <захавать> {что нужно захавать}")
			fmt.Println("---- This command let's you eat something")
			fmt.Println("")
			fmt.Println("Команда: <заказать> {что нужно заказать}")
			fmt.Println("---- This command let's you order something")
			fmt.Println("Команда: <выйти/exit> {}")
			fmt.Println("---- This command let's you exit.")
		case "exit", "выйти":
			fmt.Println("")
			fmt.Println("exiting...")
			return
		default:
			fmt.Println("Not a command, try: help")
		}

	}

}
