package main

import "fmt"

func recebeQuit(canal chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido:", <-canal)
	}
	quit <- 0
}

func enviaParaCanal(canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
		case canal <- qualquercoisa:
			qualquercoisa++
		case <-quit:
			return
		}
	}
}

func main() {
	canal := make(chan int)
	quit := make(chan int)
	go recebeQuit(canal, quit)
	enviaParaCanal(canal, quit)
}
