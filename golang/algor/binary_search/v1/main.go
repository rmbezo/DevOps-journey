package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Binary search Vs Simple search!")

	// Creating a slice

	slice := make([]int, 0, 1000000)
	// randIterations := rand.IntN(1000)
	for i := 0; i < 1000000; /* randIterations */ i++ {
		slice = append(slice, i)
	}

	// Guess logic
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Type your guess number (%v is the highest number): ", len(slice)-1)

	if ok := scanner.Scan(); !ok {
		fmt.Println("Exiting")
		return
	}

	scannerText, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error! Not an integer!")
		return
	}

	// Binary vs Simple
	simple_search(slice, scannerText)
	binary_search(slice, scannerText)
}

func simple_search(l []int, g int) int {
	simple_time := time.Now()
	for i := 0; i < len(l); i++ {
		if g == l[i] {
			fmt.Println("Simple iterations: ", i)
			fmt.Println("Time for simple search: ", time.Since(simple_time))
			return i
		}
	}
	return -1
}

func binary_search(l []int, g int) int {
	binary_time := time.Now()
	high, low := len(l)-1, 0
	i := 0
	for low <= high {
		i += 1
		mid := low + (high-low)/2
		if g == l[mid] {
			fmt.Println("Binary iterations: ", i)
			fmt.Println("Time for binary search: ", time.Since(binary_time))
			return mid
		} else if g > l[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
