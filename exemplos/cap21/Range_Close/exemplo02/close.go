package main

import "fmt"

func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}

func main() {
	c := make(chan int)
	go meuloop(10, c)
	for v := range c {
		fmt.Println("Recebido do canal:", v)
	}
}
