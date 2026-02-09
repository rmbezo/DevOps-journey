package main

import "fmt"

func main() {
	tickets_num := 123
	event_date := "12 January"
	var answer string
	fmt.Println("Do you have a ticket for a event(yes/no)?: ")
	fmt.Scan(&answer)

	var enter_wait string
	switch answer {
	case "y", "yes":
		fmt.Println("The ivents has:", tickets_num, "tickets")
		fmt.Println("The date of event:", event_date)
		fmt.Println("Press Enter, to quit...")
		fmt.Scanln(&enter_wait)
		fmt.Println("Quiting.")
	case "n", "no":
		fmt.Println("We are sorry, but you cannot acces the information.")
		fmt.Println("Press Enter, to quit...")
		fmt.Scanln(&enter_wait)
		fmt.Println("Quiting.")
	}
}

func loltest() {
	var lolmeme = 123
	fmt.Println(lolmeme)

}
