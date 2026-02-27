package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int64 = 2, 42
	if _, err := fmt.Fscan(in, &a, &b); err != nil {
		return
	}

	// path will store numbers while going backwards: b -> ... -> a
	path := make([]int64, 0, 64)
	path = append(path, b)

	for b > a {
		if b%10 == 1 {
			b = (b - 1) / 10
			path = append(path, b)
		} else if b%2 == 0 {
			b = b / 2
			path = append(path, b)
		} else {
			fmt.Println("NO")
			return
		}
	}

	if b != a {
		fmt.Println("NO")
		return
	}

	// reverse path to get a -> ... -> b
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	fmt.Println("YES")
	fmt.Println(len(path))
	for i, v := range path {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
