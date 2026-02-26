package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	os.WriteFile("/proc/self/oom_score_adj", []byte("1000"), 0644)
	fmt.Println("Aggressive memory allocation - OOM incoming!")

	var memoryHog [][]byte
	allocSize := 100 * 1024 * 1024 // 100 MB per iteration

	for i := 0; ; i++ {
		buffer := make([]byte, allocSize)
		// Fill to prevent compiler optimization
		for j := range buffer {
			buffer[j] = 1
		}
		memoryHog = append(memoryHog, buffer)

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Allocated: %d MB\n", m.Alloc/1024/1024)
	}
}
