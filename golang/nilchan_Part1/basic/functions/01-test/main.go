package main

import (
	"fmt"
	"time"
)

func main() {
	weareslow()
}

func weareslow() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d. We are so slooooow\n", i)
		time.Sleep(1 * time.Second)
	}
}
