package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	zezinho := pessoa{
		"Zezinho",
		"Carvalho",
		25,
	}
	fmt.Println(zezinho)

	mudeMe(&zezinho)
	fmt.Println(zezinho)
}

func mudeMe(p *pessoa) {
	(*p).nome = "Joãozinho" //maneira tradicional
	p.sobrenome = "Videira" //atalho
}
