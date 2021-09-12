package main

import "fmt"

func main() {
	fmt.Println(Soma(2, 2))
	fmt.Println(Subtracao(1, 1))
	fmt.Println(Multiplicacao(2, 5))
}

// func Soma retorna a soma de vários números passados como parâmetro
func Soma(n ...int) int {
	res := 0
	for _, v := range n {
		res += v
	}
	return res
}

// func Subtracao retorna a subtração de dois números passados como parâmetro
func Subtracao(n1, n2 int) int {
	return n1 - n2
}

// func Multiplicacao retorna a multiplicação de vários números passados como parâmetro
func Multiplicacao(n ...int) int {
	res := 1
	for _, v := range n {
		res *= v
	}
	return res
}
