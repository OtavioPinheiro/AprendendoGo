// Demonstre o comma ok idiom
// Link: https://play.golang.org/p/YHOMV9NYKK

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() { // preciso da go func para que eu possa colocar o valor no canal de maneira concorrente. Caso contrário terei o erro: All goroutines are asleep - deadlock
		c <- 17
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
