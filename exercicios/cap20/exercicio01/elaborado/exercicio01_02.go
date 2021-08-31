package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	criaGoroutines(2)
	wg.Wait()
}

func criaGoroutines(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func(x int) {
			fmt.Println("Goroutine nº", x+1)
			wg.Done()
		}(i)
	}
}
