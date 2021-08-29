package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var contador int64
	contador = 0
	totaldegoroutine := 100

	var wg sync.WaitGroup
	wg.Add(totaldegoroutine)

	for i := 0; i < totaldegoroutine; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Contador:", contador)
}
