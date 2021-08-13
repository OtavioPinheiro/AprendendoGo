package main

import "fmt"

func main() {
	fmt.Println(fatorialRecursivo(6))
	fmt.Println(fatorialLoops(6))
}

func fatorialRecursivo(x int) int {
	if x == 1 || x == 0 {
		x := 1
		return x
	}
	return x * fatorialRecursivo(x-1)
}

func fatorialLoops(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}
