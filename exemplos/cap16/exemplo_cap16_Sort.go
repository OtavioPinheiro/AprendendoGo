package main

import (
	"fmt"
	"sort"
)

func main() {
	power_rangers := []string{
		"Vermelho",
		"Amarela",
		"Rosa",
		"Azul",
		"Preto",
		"Verde",
		"Branco",
	}
	fmt.Println(power_rangers)
	sort.Strings(power_rangers)
	fmt.Println(power_rangers)
}
