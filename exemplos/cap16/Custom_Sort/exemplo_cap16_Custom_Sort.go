package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int           { return len(x) }
func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }
func (x ordenarPorPotencia) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	carros := []carro{
		{"Uno", 50, 5},
		{"Fusca", 20, 4},
	}
	fmt.Println(carros)
	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println(carros)
}
