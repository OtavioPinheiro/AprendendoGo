package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	trabalhadoresDiponiveis := 5
	qtdeTrabalho := 20

	go demanda(qtdeTrabalho, canal1)
	go fornecedora(trabalhadoresDiponiveis, canal1, canal2)

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

func fornecedora(disponivel int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	for i := 0; i < disponivel; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho(num int) int {
	time.Sleep(time.Second)
	return num
}
