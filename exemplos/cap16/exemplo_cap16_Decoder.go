package main

import (
	"encoding/json"
	"os"
)

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {
	imperador := Pessoa{
		Nome:          "Otávio",
		Sobrenome:     "Augusto",
		Idade:         20,
		Profissao:     "1º Imperador de Roma",
		Contabancaria: 200000.75,
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(imperador)
}
