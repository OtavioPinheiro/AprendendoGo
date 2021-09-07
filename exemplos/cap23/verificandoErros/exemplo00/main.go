package main

import "fmt"

func main() {
	n, err := fmt.Println("Olá")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
