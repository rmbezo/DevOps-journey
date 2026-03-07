package main

import (
	"fmt"
	"time"
)

// Рабочий идет в шахту
func mine(transferPoint chan int, n int) {
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахтму номер", n, "закончился")

	// channel
	transferPoint <- 10
}

func main() {
	coal := 0

	transferPoint := make(chan int)

	before := time.Now()

	for i := 0; i < 10_000; i++ {
		go mine(transferPoint, i)
	}

	time.Sleep(5 * time.Second)

	for i := 0; i < 10_000; i++ {
		coal += <-transferPoint
	}

	fmt.Println("Всего добыто угля:", coal)
	fmt.Println(time.Since(before))

}
