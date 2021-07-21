package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	meumapa := make(map[string]pessoa)

	pessoa1 := pessoa{
		nome:      "Renata",
		sobrenome: "Ingrata",
		sabores:   []string{"Chocolate", "Romeu e Julieta", "Morango", "Baunilha"},
	}
	pessoa2 := pessoa{"Otávio", "Pinetree", []string{"Chocolate", "Morango", "Maracujá", "Uva"}}

	meumapa[pessoa1.sobrenome] = pessoa1
	meumapa[pessoa2.sobrenome] = pessoa2

	for _, valor := range meumapa {
		fmt.Println("Meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.sabores {
			fmt.Println("-", valor)
		}
		fmt.Println()
	}

}
