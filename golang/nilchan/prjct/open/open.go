package main

import (
	"fmt"
	"time"
)

func main() {
	opened := false
	closed := true

	var choice string
	for {
		fmt.Println("1 - open window")
		fmt.Println("2 - close window")
		fmt.Println("3 - close program or press ENTER")
		fmt.Println("choose option: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1", "yes", "open":
			if opened == true {
				fmt.Println("Already opened")
			} else {
				openw()
				opened = true
			}
		case "2", "no", "close":
			if closed == true {
				fmt.Println("Already closed")
			} else {
				closew()
				closed = true
			}
		case "3", "exit", "":
			return
		default:
			fmt.Println("not an operator")
		}
	}
}

func openw() {
	fmt.Println("Starting opening window")
	time.Sleep(1 * time.Second)
	fmt.Println("Window now is opened")
	time.Sleep(2 * time.Second)
}

func closew() {
	fmt.Println("Starting closing window")
	time.Sleep(1 * time.Second)
	fmt.Println("Window now is closed")
	time.Sleep(2 * time.Second)
}
