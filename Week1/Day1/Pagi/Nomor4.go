package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	buffered := make(chan int, 5)
	start := time.Now()
	go Nomor4(buffered)
	for angka := range buffered {
		fmt.Println(angka)
	}
	elapsed := time.Since(start)
	fmt.Printf("Waktu Eksekusi %s\n", elapsed)
	PrintMemUsage()
	//Tidak terlihat perbedaan yang signifikan antara buffered dan unbuffered pada penggunaan di atas.
}

func Nomor4(channel chan<- int) {
	for i := 1; i <= 10; i++ {
		channel <- i
	}
	defer close(channel)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
