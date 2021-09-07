package main

import "fmt"

func main() {
	var resp1, resp2, resp3 string

	fmt.Print("Nome: ")
	_, err1 := fmt.Scan(&resp1)
	if err1 != nil {
		panic(err1)
	}

	fmt.Print("Comida favorita: ")
	_, err2 := fmt.Scan(&resp2)
	if err2 != nil {
		panic(err2)
	}

	fmt.Print("Esporte favorito: ")
	_, err3 := fmt.Scan(&resp3)
	if err3 != nil {
		panic(err3)
	}

	fmt.Println(resp1, resp2, resp3)
}
