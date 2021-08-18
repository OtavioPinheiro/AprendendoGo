package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) cumprimentar() {
	fmt.Println(p.nome, "disse: Bom dia!")
}

func main() {
	jose := pessoa{"José", 30}
	jose.cumprimentar()
}
