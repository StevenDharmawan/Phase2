package main

import (
	"fmt"
	"time"
)

func main() {
	odd := make(chan int)
	even := make(chan int)
	go oddLoop(odd)

	go evenLoop(even)
	counter := 1
	for {
		if counter == 20 {
			break
		}
		select {
		case data := <-odd:
			fmt.Println("Received an odd number: ", data)
		case data := <-even:
			fmt.Println("Received an even number: ", data)
		}
		counter++
	}
}

func oddLoop(channel chan<- int) {
	defer close(channel)
	for i := 1; i < 20; i += 2 {
		channel <- i
	}
	time.Sleep(2 * time.Second)
}

func evenLoop(channel chan<- int) {
	defer close(channel)
	for i := 2; i <= 20; i += 2 {
		channel <- i
	}
	time.Sleep(2 * time.Second)
}
