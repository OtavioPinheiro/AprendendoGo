package main

import (
	"fmt"
)

func main() {
	//for usado como while
	x := 10
	for x >= 0 {
		fmt.Printf("x = %d\n", x)
		x--
	}

	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i, "é par")
	}

	for i := 33; i <= 122; i++ {
		fmt.Printf("%d -> %v\n", i, string(i))
	}
}
