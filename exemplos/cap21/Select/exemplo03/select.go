package main

import "fmt"

func send(par chan int, impar chan int, quit chan bool) {
	numeros := 300
	for i := 0; i <= numeros; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func receive(par chan int, impar chan int, quit chan bool) {
	for {
		select {
		case num := <-par:
			fmt.Println("O número", num, "é par")
		case num := <-impar:
			fmt.Println("O número", num, "é ímpar")
		case <-quit:
			return
		}
	}
}

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)
	go send(par, impar, quit)
	receive(par, impar, quit)
}
