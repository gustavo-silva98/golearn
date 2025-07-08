package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Printa a mensagem de saudação
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
