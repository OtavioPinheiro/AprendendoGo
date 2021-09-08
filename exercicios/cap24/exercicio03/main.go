// Crie um struct "erroEspecial" que implemente a interface builtin.error
// Crie uma função que tenha um valor do tipo error como parâmetro
// Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior
package main

import "fmt"

type erroEspecial struct {
	err string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("Um erro especial acaba de ocorrer: %v", e.err)
}

func erroComoParametro(e error) {
	fmt.Println("Função que recebe um valor do tipo error como parâmetro:", e.(erroEspecial).err, "\nMétodo Error:", e)
}

func main() {
	es := erroEspecial{
		"passando por um erro especial",
	}
	erroComoParametro(es)
}
