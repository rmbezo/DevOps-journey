package main

import (
	"fmt"
	"time"
)

type Channel struct {
	Privet string
	Error  error
}

func privet(name string, c chan Channel) {
	str := "Privet" + name
	c <- Channel{str, nil}
}

func main() {
	channel := make(chan Channel)

	go privet("Gosha", channel)

	result := <-channel

	time.Sleep(100 * time.Millisecond)
	fmt.Println(result)
}
