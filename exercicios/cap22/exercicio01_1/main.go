//Resolva o problema do código.
//all goroutines are asleep - deadlock
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
