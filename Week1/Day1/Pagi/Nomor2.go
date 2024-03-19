package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		mutex.Lock()
		for i := 'a'; i <= 'j'; i++ {
			fmt.Printf("%c\n", i)
		}
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
		mutex.Unlock()
	}()
	wg.Wait()
}
