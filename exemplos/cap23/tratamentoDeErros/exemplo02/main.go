package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("./exemplos/cap23/tratamentoDeErros/exemplo02/log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("./exemplos/cap23/tratamentoDeErros/exemplo02/no-file.txt")
	if err != nil {
		log.Println("Erro ocorrido:", err)
	}
	defer f2.Close()

	fmt.Println("Verifique o arquivo log.txt nesta pasta")
}
