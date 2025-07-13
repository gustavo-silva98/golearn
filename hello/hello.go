package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	// Chama a mensagem de greetings 
	message,err := greetings.Hello("")
	// Se retorna um erro, printa no console  e exit o programa

	if err != nil {
		log.Fatal(err)
	}

	// Se nenhum erro for retornado, printa a mensagem de greeting
	fmt.Println(message)
}
