package main

import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	messageChan1 := make(chan Message)
	messageChan2 := make(chan Message)

	go func() {
		for {
			messageChan1 <- Message{
				Author: "Friend 1",
				Text:   "Privet",
			}
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			fmt.Println("Время отправки сообщения вторым другом: ", time.Now())
			messageChan2 <- Message{
				Author: "Friend 2",
				Text:   "Kak dela?",
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg1 := <-messageChan1:
			fmt.Println("Я получил сообщение от:", msg1.Author, "текст сообщения:", msg1.Text)
		case msg2 := <-messageChan2:
			fmt.Println("Я получил сообщение от:", msg2.Author, "текст сообщения:", msg2.Text, "Время получения сообщения от второго друга:", time.Now())
		}

	}
}
