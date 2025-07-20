package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice_sum := scanner()
	sum := sum_slice(slice_sum)
	fmt.Printf("O valor da soma de %v é %v\n ", slice_sum, sum)
}

func sum_slice(slice []int) int {

	sum := 0
	for _, value := range slice {
		sum = sum + value
	}
	return sum
}

func scanner() []int {
	var slice_resp []int
	fmt.Println("Digite o número para a soma: ")
	fmt.Println("Caso queira finalizar o ciclo, digite F.")

	// Loop Infinito para adicionar números numa lista/slice
	for {
		scanner := bufio.NewReader(os.Stdin)       // linha que lê a entrada do usuario
		numero_str, _ := scanner.ReadString('\n')  // lê a entrada do usuário até encontrar newline
		numero_str = strings.TrimSpace(numero_str) // Tira os espaços ao redor do String

		if numero_str == "F" || numero_str == "f" {
			break
		} // Se a entrada for f ou F, ele vai parar o loop

		numero, err := strconv.Atoi(numero_str) // Tenta converter o número string para Int
		if err != nil {
			fmt.Println("Digite um número inteiro, por favor.")
			continue
		} // Caso dê erro, retorna a msg e pula a vez do for
		slice_resp = append(slice_resp, numero) // adiciona entrada do usuario a lista de numeros
	}
	return slice_resp // Retorna slice de fato

}
