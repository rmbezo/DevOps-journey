package main

import "fmt"

func main() {
	// 5! = 1*2*3*4*5
	previousI := 1
	num := 10
	for i := 1; i <= num; i++ {
		previousI = i * previousI
		fmt.Println(previousI)
	}
}
