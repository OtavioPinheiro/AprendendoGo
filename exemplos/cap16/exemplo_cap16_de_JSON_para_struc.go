package main

import (
	"encoding/json"
	"fmt"
)

type informacoes struct {
	Nome          string  `json: "Nome"`
	Sobrenome     string  `json: "Sobrenome"`
	Idade         int     `json: "Idade"`
	Profissao     string  `json: "Profissao"`
	Contabancaria float64 `json: "Contabancaria"`
}

func main() {
	sb := []byte(`{
		"Nome": "Otavius",
		"Sobrenome": "Augustus",
		"Idade": 30,
		"Profissao": "1º Imperador de Roma",
		"Contabancaria": 200000000
	}`)

	var otavius informacoes
	err := json.Unmarshal(sb, &otavius)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(otavius)
	fmt.Println("Meu nome é", otavius.Nome, otavius.Sobrenome)
	fmt.Println("Eu sou o", otavius.Profissao)
}
