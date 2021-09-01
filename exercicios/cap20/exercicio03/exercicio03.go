package main

import (
	"fmt"
	"runtime"
	"sync"
)

const qtdeGoroutines = 500
var wg sync.WaitGroup

func main() {
	criaGoroutines(qtdeGoroutines)
	wg.Wait()
}
//OBS.: Se o valor total de Goroutines for diferente do total do contador,
//então temos uma race condition(condição de corrida)

func criaGoroutines(qtdeGoroutines int) {
	wg.Add(qtdeGoroutines)
	contator := 0
	for i := 0; i < qtdeGoroutines; i++ {
		go func() {
			v := contator
			runtime.Gosched()
			v++
			contator = v
			wg.Done()
		}()
	}

	fmt.Println("Total de Goroutines:\t", qtdeGoroutines)
	fmt.Println("Total do contador:\t", contator)
}