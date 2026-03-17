package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	slice := []string{"Alibaba", "Boris", "Cara", "Denis", "Eren"}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type first letter of search: ")
	if ok := scanner.Scan(); !ok {
		fmt.Println("Exiting")
		return
	} else if scanner.Text() == "" {
		fmt.Println("Exiting")
		return
	}
	guess := strings.ToLower(string(scanner.Text()))

	// fmt.Printf("%v, %T\n", guess, guess)
	// fmt.Println(slice)

	fmt.Println(binary_serch(slice, guess))
	fmt.Println(" ")
	fmt.Println(simple_search(slice, guess))
}

func simple_search(l []string, g string) int {
	for i := 0; i < len(l); i++ {
		if g == strings.ToLower(string(l[i][0])) {
			fmt.Println("Simple search iterations: ", i)
			return i
		} else {
			continue
		}
	}
	return -1
}

func binary_serch(l []string, g string) int {
	high, low := len(l)-1, 0
	i := 0
	for low <= high {
		i += 1
		mid := low + (high-low)/2
		if g == strings.ToLower(string(l[mid][0])) {
			fmt.Println("Binary search iterations: ", i)
			return mid
		} else if g > strings.ToLower(string(l[mid][0])) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
