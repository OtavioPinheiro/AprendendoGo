/*OBS.: O problema desse código não foi ressolvido com comma ok e sim simplesmente trocando a ordem de execução.
A linha quit <- true passou a ser executada antes do close(par) e close(impar)
*/
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
	quit <- true
	close(par)
	close(impar)
}

func receive(par chan int, impar chan int, quit chan bool) {
	for {
		select {
		case num := <-par:
			fmt.Println("O número", num, "é par")
		case num := <-impar:
			fmt.Println("O número", num, "é ímpar")
		case _, ok := <-quit:
			if !ok {
				fmt.Println("Comma OK:", ok)
			} else {
				fmt.Println("Comma OK:", ok)
			}
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
