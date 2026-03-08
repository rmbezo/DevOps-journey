package main

import "fmt"

func main() {

	var ch chan string

	go func() {
		ch <- "Gello"
	}()

	// go func() {
	// 	name := <-ch
	// }()
	// close(ch)

	v := <-ch

	fmt.Println(v)
}
