package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getCost(sub string, target string) int {
	cost := 0
	for i := 0; i < len(target); i++ {
		if sub[i] != target[i] {
			cost++
		}
	}
	return cost
}

func solve() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 100005)
	sc.Buffer(buf, 100005)

	if !sc.Scan() {
		return
	}
	s := sc.Text()
	n := len(s)

	tbank := "tbank"
	study := "study"
	m := 5

	if n < 10 {
		fmt.Println(-1)
		return
	}

	costT := make([]int, n-m+1)
	costS := make([]int, n-m+1)

	for i := 0; i <= n-m; i++ {
		costT[i] = getCost(s[i:i+m], tbank)
		costS[i] = getCost(s[i:i+m], study)
	}

	sufMinS := make([]int, n-m+1)
	sufMinT := make([]int, n-m+1)

	sufMinS[n-m] = costS[n-m]
	sufMinT[n-m] = costT[n-m]

	for i := n - m - 1; i >= 0; i-- {
		sufMinS[i] = min(costS[i], sufMinS[i+1])
		sufMinT[i] = min(costT[i], sufMinT[i+1])
	}

	ans := n

	for i := 0; i <= n-2*m; i++ {
		ans = min(ans, costT[i]+sufMinS[i+m])
		ans = min(ans, costS[i]+sufMinT[i+m])
	}

	fmt.Println(ans)
}

func main() {
	solve()
}
