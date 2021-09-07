package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("./exemplos/cap23/verificandoErros/exemplo02/nomes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err2 := ioutil.ReadAll(f)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println(string(bs))
}
