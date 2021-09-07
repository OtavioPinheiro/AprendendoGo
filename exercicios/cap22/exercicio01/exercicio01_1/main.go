// Faça o código funcionar
// Link: https://play.golang.org/p/j-EA6003P0
// Erro: all goroutines are asleep - deadlock
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}
