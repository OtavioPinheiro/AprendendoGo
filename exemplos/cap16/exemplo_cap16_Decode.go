package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
	[
		{"Nome": "Zé", "Texto": "Knock knock."},
		{"Nome": "Jão", "Texto": "Who's there?"},
		{"Nome": "Zé", "Texto": "Go fmt."},
		{"Nome": "Jão", "Texto": "Go fmt who?"},
		{"Nome": "Zé", "Texto": "Go fmt yourself!"}
	]
`
	type Message struct {
		Nome, Texto string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Nome, m.Texto)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}
