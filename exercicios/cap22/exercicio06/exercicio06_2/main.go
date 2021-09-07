// Escreva um programa que coloque 100 números em um canal, retire-os do canal e demonstre-os
package main

import "fmt"

func main() {
	c := make(chan int)

	go numeros(c)

	for i := range c {
		fmt.Println(i)
	}
}

func numeros(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
