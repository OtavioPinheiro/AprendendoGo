package main

import "fmt"

func Media(v1 []float64) float64 {
	return (v1[0] + v1[1]) / float64(len(v1))
}

func main() {
	fmt.Println("Testando...")
}
