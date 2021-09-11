package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "20julho1980"
	senhaErrada := "senhaErrada77"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaErrada)))
}
