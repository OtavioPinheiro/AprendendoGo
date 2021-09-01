//Resolvendo a condição de corrida do exercício 03 usando Mutex
package main

import (
	"fmt"
	"runtime"
	"sync"
)

const qtdeGoroutines = 500
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	criaGoroutines(qtdeGoroutines)
	wg.Wait()
}

func criaGoroutines(qtdeGoroutines int) {
	wg.Add(qtdeGoroutines)
	contator := 0
	for i := 0; i < qtdeGoroutines; i++ {
		go func() {
			mu.Lock()
			v := contator
			runtime.Gosched()
			v++
			contator = v
			mu.Unlock()
			wg.Done()
		}()
	}

	fmt.Println("Total de Goroutines:\t", qtdeGoroutines)
	fmt.Println("Total do contador:\t\t", contator)
}