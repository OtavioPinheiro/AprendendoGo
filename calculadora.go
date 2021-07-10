package main

import (
	"fmt"
)

func main() {
	var operacao string
	var resultado int
	x := 10
	y := 20

	operacao = "salario"

	switch operacao {
	case "soma":
		resultado = x + y
		fmt.Printf("Resultado: %d", resultado)
	case "subtracao":
		resultado = y - x
		fmt.Printf("Resultado: %d", resultado)
	case "salario":
		resultado = x + y
		fmt.Printf("Salário: R$%d,00\n", resultado)
		fallthrough
	case "despesa":
		resultado = resultado - y
		fmt.Printf("Conta bancária: R$%d,00", resultado)
	default:
		fmt.Println("Nenhuma operação informada!")
	}

}
