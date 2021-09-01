package main

import "fmt"

type pessoa struct{
	nome string
	idade int
}

func (p *pessoa ) falar()  {
	fmt.Println(p.nome, "diz oi!")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	p1 := pessoa{
		"Otávio Augusto",
		200,
	}
	p1.falar() //essa maneira de executar esse método é um atalho! A maneira "correta" seria (&p1).falar()
	dizerAlgumaCoisa(&p1) //para essa função não podemos passar apenas p1, temos que passar o endereço de memória de p1
	//p1, pelo fato de implementar o método falar(), faz parte da interface humanos, automaticamente.
}
