package main

import (
	"fmt"
)

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"d", "e", "f"}
	c := append(a, b...)
	d := []int{1, 2, 3}
	// e := append(a, d...)
	fmt.Println(c)
	fmt.Println(d)
}
