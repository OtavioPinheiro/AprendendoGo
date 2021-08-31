package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		fmt.Println("Função 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Função 2")
		wg.Done()
	}()
	wg.Wait()
}
