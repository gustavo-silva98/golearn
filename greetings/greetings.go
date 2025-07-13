package greetings

import (
	"fmt"
	"errors"
)

// Hello retorna uma saudação para cada pessoa nomeada

func Hello(name string) (string,error) {
	// Retorna um nome interpolado com uma string
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Opa!, Tudo certo, %v?", name)
	return message, nil
}
