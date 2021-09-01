//Resolvendo a condição de corrida do exercício 03 usando atomic
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

const qtdeGoroutines = 500
var wg sync.WaitGroup

func main() {
	criaGoroutines(qtdeGoroutines)
	wg.Wait()
}

func criaGoroutines(qtdeGoroutines int) {
	wg.Add(qtdeGoroutines)
	var contator int64 = 0
	for i := 0; i < qtdeGoroutines; i++ {
		go func() {
			atomic.AddInt64(&contator, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}

	fmt.Println("Total de Goroutines:\t", qtdeGoroutines)
	fmt.Println("Total do contador:\t\t", contator)
}