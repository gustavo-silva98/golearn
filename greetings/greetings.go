package greetings

import "fmt"

// Hello retorna uma saudação para cada pessoa nomeada

func Hello(name string) string {
	// Retorna um nome interpolado com uma string
	message := fmt.Sprintf("Opa!, Tudo certo, %v?", name)
	return message
}
