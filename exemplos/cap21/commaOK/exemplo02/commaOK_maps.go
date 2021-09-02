package main

import "fmt"

func main() {
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	var val1 = timeZone["PST"]
	var val2 = timeZone["invalid"]
	var val3, ok1 = timeZone["UTC"]
	var val4, ok2 = timeZone["invalid"]

	fmt.Println("> Sem comma ok")
	fmt.Printf("val1: %d\tval2: %d\n", val1, val2)
	fmt.Println("\n> Com comma ok")
	fmt.Printf("val3: %d commaOK: %t\tval4: %d commaOK: %t\n", val3, ok1, val4, ok2)
}
