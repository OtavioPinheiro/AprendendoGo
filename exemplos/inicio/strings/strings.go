package main

import (
	"fmt"
)

func main() {
	s := "ascii éùüàáó \u9999"
	ola := "Olá GO"
	olaSlice := []byte(s)

	for _, v := range olaSlice {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", v, v, v, v, v)
	}

	fmt.Println()

	for _, v := range s {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", v, v, v, v, v)
	}

	fmt.Println()

	for i := 0; i < len(ola); i++ {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", ola[i], ola[i], ola[i], ola[i], ola[i])
	}

	fmt.Println()

	for _, i := range ola {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", i, i, i, i, i)
	}
}
