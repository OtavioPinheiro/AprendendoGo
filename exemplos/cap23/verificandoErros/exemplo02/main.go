package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("./exemplos/cap23/verificandoErros/exemplo02/nomes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	defer fmt.Println("Rodou a função que estava em defer")

	r := strings.NewReader(fmt.Sprintf("Sherlock Holmes"))

	io.Copy(f, r)
}
