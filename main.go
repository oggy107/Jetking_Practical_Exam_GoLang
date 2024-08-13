package main

import "fmt"

func sum(arr []int, c chan int) {
	s := 0

	for _, v := range arr {
		s += v
	}

	c <- s
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	channel := make(chan int)

	// Split the array into two parts and pass them to two separate go routines
	go sum(a[:len(a)/2], channel)
	go sum(a[len(a)/2:], channel)

	// Receive the two sums from the two separate go routines
	x, y := <-channel, <-channel

	fmt.Println("The sum returned by first go routine is:", x)
	fmt.Println("The sum returned by second go routine is:", y)
	fmt.Println("The sum is:", x+y)
}
