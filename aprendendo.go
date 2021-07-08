package main

import (
	"fmt"
)

var x int = 10
var y float32 = 12.2

func main() {
	txt := "Eu sou "
	cantando := string(0x266c) //Recomenda-se usar o fmt.Sprint()
	var rsp string
	z := 3.14153111
	ok := true
	nome := "Teste"
	fmt.Println("Aprendendo GO!")
	fmt.Println("|   Valor   |    Tipo   |")
	fmt.Println("|-----------|-----------|")
	fmt.Printf("|  %v  |  %T  |\n", x, x)
	fmt.Printf("|  %v  |  %T  |\n", y, y)
	fmt.Printf("|  %v  |  %T  |\n", z, z)
	fmt.Printf("|  %v  |  %T  |\n", ok, ok)
	fmt.Printf("|  %v  |  %T  |\n\n", nome, nome)

	rsp = txt + fmt.Sprint(10)
	fmt.Println(rsp)
	fmt.Println("Lá lá lá " + cantando)
}
