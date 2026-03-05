package main

import "fmt"

func main() {
	weather := make(map[int]int)

	AddData(weather)

	fmt.Println(weather)
}

func AddData(m map[int]int) {
	for {
		var choice int
		fmt.Println("Hello, this is Admin's CLI")
		fmt.Println("Choose the action:")
		fmt.Println("1-Add new data, 2-Quit")
		fmt.Print("Action: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var date, temperature int
			fmt.Println("Choose: Updating the data.")
			fmt.Print("Enter date: ")
			fmt.Scan(&date)
			if date > 31 || date < 1 {
				fmt.Println("Not right type of date")
				break
			}
			fmt.Print("Enter temperature: ")
			fmt.Scan(&temperature)
			m[date] = temperature
			fmt.Println("Data was added.")
		case 2:
			return
		}
	}
}
