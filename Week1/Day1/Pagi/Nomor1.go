package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 'a'; i <= 'j'; i++ {
		go func(input int32) {
			fmt.Printf("%c\n", input)
		}(i)
	}
	for i := 1; i <= 10; i++ {
		go func(input int) {
			fmt.Println(input)
		}(i)
	}
	time.Sleep(time.Second)
}
