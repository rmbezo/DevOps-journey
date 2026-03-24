package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD int64 = 1_000_000_007

func colorWays(lengths []int, maxK int) []int64 {
	sort.Ints(lengths)

	dp := make([]int64, maxK+1)
	dp[0] = 1

	for _, L := range lengths {
		for j := maxK; j >= 1; j-- {
			available := L - (j - 1)
			if available > 0 {
				dp[j] = (dp[j] + dp[j-1]*int64(available)) % MOD
			}
		}
	}

	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	if n == 1 {
		if k == 1 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		return
	}

	maxBishops := 2*n - 2
	if k > maxBishops {
		fmt.Println(0)
		return
	}

	var color1, color2 []int
	for d := 0; d < 2*n-1; d++ {
		length := n - abs((n-1)-d)
		if d%2 == 0 {
			color1 = append(color1, length)
		} else {
			color2 = append(color2, length)
		}
	}

	limit := min(k, maxBishops)

	ways1 := colorWays(color1, limit)
	ways2 := colorWays(color2, limit)

	var ans int64
	for i := 0; i <= k; i++ {
		if i < len(ways1) && k-i < len(ways2) {
			ans = (ans + ways1[i]*ways2[k-i]) % MOD
		}
	}

	fmt.Println(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
