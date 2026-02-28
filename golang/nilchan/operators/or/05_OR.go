package main

import "fmt"

func main() {
	icec := false
	pcclub := true

	if icec || pcclub {
		fmt.Println("Okay, let's go for a walk!")
	} else {
		fmt.Println("No i dont want to go for a walk..")
	}

}
