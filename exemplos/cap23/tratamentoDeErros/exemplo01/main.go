package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("./exemplos/cap23/tratamentoDeErros/exemplo/no-file.txt")
	if err != nil {
		fmt.Println("Erro ocorrido:", err)
		log.Println("Erro ocorrido:", err)
		panic(err)
		log.Fatalln(err)
	}
}
