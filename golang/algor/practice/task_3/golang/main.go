package main

import (
	"bufio"
	"fmt"
	"os"
)

func solveOne(s string) int64 {
	n := len(s)

	allOnes := true
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			allOnes = false
			break
		}
	}
	if allOnes {
		return int64(n) * int64(n)
	}

	ss := s + s
	cur := 0
	best := 0

	for i := 0; i < len(ss); i++ {
		if ss[i] == '1' {
			cur++
			if cur > n {
				cur = n
			}
			if cur > best {
				best = cur
			}
		} else {
			cur = 0
		}
	}

	sum := best + 1
	h := sum / 2
	w := sum - h

	return int64(h) * int64(w)
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var s string
		fmt.Fscan(in, &s)
		fmt.Fprintln(out, solveOne(s))
	}
}
