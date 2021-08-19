package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	ContaBancaria float64
}

func main() {
	p1 := Pessoa{
		"Carlos",
		"César",
		27,
		"Empresário",
		10500,
	}

	p2 := Pessoa{
		"Otavio",
		"Augusto",
		28,
		"1º Imperador de Roma",
		200000,
	}

	a, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(a))
	fmt.Println(string(b))
}
