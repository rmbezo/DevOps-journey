package main

import (
	"fmt"
	"strconv"
)

func main() {

	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		for i := 1; i < 1_000; i++ {
			intCh <- i
		}
	}()

	go func() {
		for i := 1; i < 1_000; i++ {
			strCh <- "hi" + strconv.Itoa(i)
		}
	}()

	for {
		select {
		case number := <-intCh:
			fmt.Println(number)
		case str := <-strCh:
			fmt.Println(str)
		}
	}
}
