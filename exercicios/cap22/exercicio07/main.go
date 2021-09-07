// Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
// Tire estes números do canal e demonstre-os;

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	numeros := 100
	qtdeGoroutines := 10

	go geraGoroutines(c, numeros, qtdeGoroutines)

	for i := 0; i < numeros; i++ {
		fmt.Println(i, "\t", <-c)
	}
}

func geraGoroutines(c chan int, numeros, qtdeGoroutines int) {

	for i := 0; i < qtdeGoroutines; i++ {
		go func(n int) {
			for j := 0; j < n; j++ {
				c <- j
			}
		}(numeros / 10)
	}
}
