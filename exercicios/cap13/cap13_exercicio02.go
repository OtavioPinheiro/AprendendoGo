package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}

	fmt.Println(função1(a...))
	fmt.Println(função2(b))
}

func função1(x ...int) int {
	somatorio := 0
	for _, v := range x {
		somatorio += v
	}
	return somatorio
}

func função2(x []int) int {
	somatorio := 0
	for _, v := range x {
		somatorio += v
	}
	return somatorio
}
