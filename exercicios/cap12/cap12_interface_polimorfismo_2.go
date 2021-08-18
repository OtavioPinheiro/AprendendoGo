package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type retangulo struct {
	comprimento, altura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.comprimento * r.altura
}

func (r retangulo) perimetro() float64 {
	return 2*r.comprimento + 2*r.altura
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func medida(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	r := retangulo{comprimento: 3, altura: 4}
	c := circulo{raio: 5}

	medida(r)
	medida(c)
}
