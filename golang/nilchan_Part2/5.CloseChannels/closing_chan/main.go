package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// type User struct {
// 	Name string
// 	Age  int
// }

func main() {
	trasferPoint := make(chan int)

	//miner
	go func() {
		// trasferPoint <- 3 + rand.IntN(4)
		iterations := 3 + rand.IntN(4)
		for i := 0; i < iterations; i++ {
			trasferPoint <- 10
			time.Sleep(300 * time.Millisecond)
		}
		close(trasferPoint)
	}()

	coal := 0
	// for {
	// 	v, ok := <-trasferPoint
	// 	if !ok {
	// 		fmt.Println("Шахтер устал. Уходит домой...")
	// 		break
	// 	}
	// 	coal += v

	// 	fmt.Println(coal)

	// }

	for v := range trasferPoint {
		coal += v
		fmt.Println(coal)
	}

	// slice := []User{}

	// element := User{
	// 	Name: "Gosha",
	// 	Age:  12,
	// }

	// slice = append(slice, element)

	// fmt.Printf("Name: %s, Age: %v\n", slice[0].Name, slice[0].Age)
}
