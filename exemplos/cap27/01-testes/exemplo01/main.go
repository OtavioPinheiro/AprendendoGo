package main

import "fmt"

func adicao(i ...int) int {
	resultado := 0
	for _, v := range i {
		resultado += v
	}
	return resultado
}

func main() {
	res := adicao(1, 5)
	fmt.Println(res)
}
