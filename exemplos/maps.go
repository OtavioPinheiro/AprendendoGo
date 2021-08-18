package main

import "fmt"

func main() {
	amigos := map[string]int{
		"Charles": 998796958,
		"Rebecca": 99995845,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["Rebecca"])

	//Adicionando um elemento novo
	amigos["Jéssica"] = 78898695

	fmt.Println(amigos)
	fmt.Println(amigos["Jéssica"], "\n")

	//comma ok idiom
	if será, ok := amigos["Gasparzinho"]; !ok {
		fmt.Println("Esse nome não consta na lista de amigos D=")
	} else {
		fmt.Println(será, ok)
	}
}
