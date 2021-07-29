package main

import "fmt"

func main() {
	defer fmt.Println("der Anfang")
	defer fmt.Println("das Ende ist")
	defer fmt.Println("das Ende und")
	defer fmt.Println("Der Anfang ist")
}
