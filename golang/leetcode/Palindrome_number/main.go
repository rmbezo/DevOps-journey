package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// Converting integer
	stringX := strconv.Itoa(x)

	isPalindrome := true
	for i := 0; i < len(stringX)/2+1; i++ {
		if stringX[i] != stringX[len(stringX)-1-i] {
			isPalindrome = false
		}
	}
	return isPalindrome
}
