package main

import (
	"context"
	"fmt"
	"time"
)

// Parent context
func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():

			// ...
			fmt.Println("Foo ended")
			return
		default:
			fmt.Println("Foo")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// children context
func boo(ctx context.Context) {
	for {
		select {

		// ...
		case <-ctx.Done():
			fmt.Println("Boo ended")
			return
		default:
			fmt.Println("Boo")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go foo(parentContext)
	go boo(childContext)

	time.Sleep(1 * time.Second)
	childCancel()

	time.Sleep(1 * time.Second)
	parentCancel()

}
