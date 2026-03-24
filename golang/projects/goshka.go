package main

import (
	"fmt"
	//"math/rand/v2"
	"bufio"
	"os"
	"strconv"
	"time"
)

type User struct {
	name  string
	score int
}

type UserSlice struct {
	s []User
}

// TO DO
// Step 1: Implement your Player struct and use the sort.Slice (or the newer slices.SortFunc) to sort them.
// Step 2: Wrap your Linear and Binary searches in Goroutines and use a channel to send the result back to main.
// Step 3: Write a simple Unit Test for your binarySearch function to make sure it handles edge cases (like an empty slice or a missing number).

func main() {
	fmt.Println("Hello!")
	fmt.Println("Let's play with algorithms!")

	//Scanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("First of all let's define the size of our slice")
	fmt.Print("Type your size of slice here(int only): ")
	if ok := scanner.Scan(); !ok {
		fmt.Println("Error!")
		return
	} else if scanner.Text() == "" || scanner.Text() == " " {
		fmt.Println("Exiting program....")
		return
	}

	size, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error!")
		return
	}

	// Creating slices
	sliceIntBinary := []int{}
	sliceStringBinary := []string{"a", "b", "c", "d", "e", "f", "g", "h", "q", "z"}

	for i := 0; i < size; i++ {
		sliceIntBinary = append(sliceIntBinary, i)
		sliceStringBinary = append(sliceStringBinary, string(i))
	}

	// Channel
	tranferPoint := make(chan int)
	// Binary search
	for {
		//fmt.Println("Okay let's start binary search with integers")
		fmt.Print("Type your number for search: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Error!")
			return
		} else if scanner.Text() == "" || scanner.Text() == " " {
			fmt.Println("Exiting program....")
			return
		}

		numberGuessBinary, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error!")
			return
		}

		go binarySearchInt(sliceIntBinary, numberGuessBinary, tranferPoint)
		go simpleSearchInt(sliceIntBinary, numberGuessBinary, tranferPoint)
	}

	// Bubble sort

	// Stack and Recursion

	// O(n!)

	//
}

func simpleSearchInt(l []int, g int, ch chan int) int {
	simpleStart := time.Now()
	for i := 0; i < len(l); i++ {
		if l[i] == g {
			fmt.Printf("Simple search Int iterations: %v, time: %v\n", i, time.Since(simpleStart))
			ch <- i
			return i
		}
	}
	fmt.Println("Simple search didn't found number.")
	return -1
}

func binarySearchInt(l []int, g int, ch chan int) int {
	binaryStart := time.Now()
	high, low := len(l)-1, 0
	i := 0
	for low <= high {
		i += 1
		mid := low + (high-low)/2
		if g == l[mid] {
			fmt.Printf("Binary search takes %v iterations. Time: %v\n", i, time.Since(binaryStart))
			ch <- mid
			return mid
		} else if g > l[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Not finded")
	return -1
}
