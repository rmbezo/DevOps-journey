package main

import (
	"bufio"
	"fmt"
	"os"
)

func min3(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}

func max3(a, b, c int) int {
	if a < b {
		a = b
	}
	if a < c {
		a = c
	}
	return a
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	const MAXA = 100000

	freq := make([]int, MAXA+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		freq[a[i]]++
	}

	good := make([]bool, MAXA+1)

	for i := 0; i < n; i++ {
		x := a[i]
		y := a[(i+1)%n]
		z := a[(i+2)%n]

		mn := min3(x, y, z)
		mx := max3(x, y, z)

		good[mn] = true
		good[mx] = true
	}

	for i := 0; i < n; i++ {
		val := a[i]
		ans := n - freq[val]
		if !good[val] {
			ans++
		}

		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, ans)
	}
	fmt.Fprintln(out)
}
