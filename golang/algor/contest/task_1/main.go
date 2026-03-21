package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	cnt := [10]int{}
	for _, ch := range s {
		cnt[ch-'0']++
	}

	minNonZero := 1
	for cnt[minNonZero] == 0 {
		minNonZero++
	}

	result := make([]byte, 0, len(s))

	result = append(result, byte(minNonZero)+'0')
	cnt[minNonZero]--

	for i := 0; i < cnt[0]; i++ {
		result = append(result, '0')
	}

	for d := 1; d <= 9; d++ {
		for i := 0; i < cnt[d]; i++ {
			result = append(result, byte(d)+'0')
		}
	}

	fmt.Println(string(result))
}
