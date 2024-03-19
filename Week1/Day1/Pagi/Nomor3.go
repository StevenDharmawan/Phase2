package main

import (
	"fmt"
)

func main() {
	//var wg sync.WaitGroup, &wg, wg *sync.WaitGroup
	channel := make(chan int)
	go Nomor3(channel)
	for angka := range channel {
		fmt.Println(angka)
	}
}

func Nomor3(channel chan<- int) {
	//defer wg.Done()
	for i := 1; i <= 10; i++ {
		channel <- i
	}
	defer close(channel)
}
