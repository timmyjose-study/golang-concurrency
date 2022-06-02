package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch := make(chan int)
	go sum(nums[:len(nums)/2], ch)
	go sum(nums[len(nums)/2:], ch)

	total := <-ch + <-ch
	fmt.Printf("Total = %d\n", total)

}

func sum(nums []int, ch chan<- int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	ch <- sum
}