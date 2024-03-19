package main

import (
	"fmt"
)

func main() {
	oddCh := make(chan int)
	evenCh := make(chan int)
	errorCh := make(chan int)
	go numbers(oddCh, evenCh, errorCh)

	for i := 0; i < 30; i++ {
		select {
		case even := <-evenCh:
			fmt.Printf("Received an even number\t: %v\n", even)

		case odd := <-oddCh:
			fmt.Printf("Received an odd number\t: %v\n", odd)

		case err := <-errorCh:
			fmt.Printf("Error: number %v is greater than 20\n", err)
		}
	}
}

func numbers(oddCh chan<- int, evenCh chan<- int, errorCh chan<- int) {
	for i := 1; i <= 30; i++ {
		if i > 20 {
			errorCh <- i
		} else if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	close(errorCh)
	close(oddCh)
	close(evenCh)
}
