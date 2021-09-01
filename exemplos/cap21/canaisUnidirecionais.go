package main

import "fmt"

func main() {
	canal := make(chan string)
	go enviar(canal)
	receber(canal)
}

func enviar(enviar chan <- string) {
	fmt.Println("Enviando os dados...")
	enviar <- "pacote de dados"
}

func receber(receber <- chan string) {
	fmt.Println("Recebendo os dados ...", <- receber)
}
