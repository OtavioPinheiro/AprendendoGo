package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go demanda(10, canal1)
	go fornecedora(canal1, canal2)

	for i := range canal2 {
		fmt.Println(i)
	}

}

func demanda(d int, canal chan int) {
	for i := 0; i < d; i++ {
		canal <- i
	}
	close(canal)
}

func fornecedora(canal1, canal2 chan int) {
	wg := sync.WaitGroup{}
	for i := range canal1 {
		wg.Add(1)
		go func(i int) {
			canal2 <- trabalho(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(canal2)
}

func trabalho(num int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return num
}
