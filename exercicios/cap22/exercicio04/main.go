// Faça o código funcionar
// Link: https://play.golang.org/p/MvL6uamrJP
// use um select statement para demonstrar os valores do canal

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int { // irá retornar um canal receiver, só poderei remover os valores do canal que será retornado por esta função.
	c := make(chan int) // canal bidirecional

	go func() {
		for i := 0; i < 100; i++ {
			c <- i // alimentando o canal c com inteiros de 0 a 99
		}
		close(c)
		q <- 0 // alimento o canal q com 0
	}()

	return c
}

func receive(c <-chan int, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Canal:", v)
		case <-q:
			return
		}
	}
}
