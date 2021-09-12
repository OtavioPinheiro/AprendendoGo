package main

import "fmt"

func main() {
	fmt.Println(Soma(2, 2))
	fmt.Println(Subtracao(2, 2))
}

func Soma(n ...int) int {
	res := 0
	for _, v := range n {
		res += v
	}
	return res
}

func Subtracao(n1, n2 int) int {
	res := n1 - n2
	return res
}
