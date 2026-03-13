package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	charCount := make(map[rune]int)

	fmt.Print("Type word: ")
	if ok := scanner.Scan(); !ok {
		fmt.Println("Error!")
		return
	}

	problem := scanner.Text()

	for _, char := range problem {
		charCount[char] += 1
	}

	hasUnique := false
	for _, char := range problem {
		if charCount[char] == 1 {
			fmt.Println("First unique char is:", string(char))
			hasUnique = true
			break
		}
	}

	// PALINDROME CHECK

	sliceChar := []rune(problem)

	isPalindrome := true
	for i := 0; i < len(sliceChar)/2; i++ {
		if sliceChar[i] != sliceChar[len(sliceChar)-1-i] {
			isPalindrome = false
			break
		}

	}

	if !hasUnique {
		fmt.Println("Text has no unique characters.")
	}

	if isPalindrome {
		fmt.Println("Text is palindrome")
	} else {
		fmt.Println("Text is not palindrome")
	}
}
