package main

import "math"

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

type triangulo struct {
	lado1, lado2, lado3 float64
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

func main() {

}
