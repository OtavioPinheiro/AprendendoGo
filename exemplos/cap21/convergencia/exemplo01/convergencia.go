package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go envia(par, impar)
	go recebe(par, impar, converge)

	for v := range converge {
		fmt.Println("Valor recebido:", v)
	}

}

func envia(par, impar chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n % 2 == 0 {
			par <- n
		} else {
			impar <- n
		}
	}
	close(par)
	close(impar)
}

func recebe(par, impar, converge chan int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for v := range par {
			converge <- v
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for v := range impar {
			converge <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(converge)
}
