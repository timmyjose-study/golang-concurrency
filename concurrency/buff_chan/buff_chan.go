package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// this would cause a deadlock since the call is blocking,
	// and nothing is going to ever fill that channel with a value
	//fmt.Println(<-ch)

	// this is fine though since we do fill in the channel's queue with
	// a value using the goroutine

	go func() {
		ch <- 3
	}()

	fmt.Println(<-ch)
}