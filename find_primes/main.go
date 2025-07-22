package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	num := scanner()
	start := time.Now()
	primes := make([]int, num)
	next_prime_index := 0
	if next_prime_index < num {
		primes[next_prime_index] = 2
		next_prime_index++
	}

	i := 3

	for next_prime_index < num {
		if isPrime(primes[:next_prime_index], i) {
			primes[next_prime_index] = i
			next_prime_index++
		}
		i += 2
	}
	elapsed := time.Since(start)
	fmt.Printf("Tempo percorrido: %v\n", elapsed)

}

func scanner() int {
	fmt.Println("Digite o número de primos que você quer encontrar.")
	for {
		reader := bufio.NewReader(os.Stdin)
		num_str, _ := reader.ReadString('\n')
		num_str = strings.TrimSpace(num_str)
		num, err := strconv.Atoi(num_str)
		if err != nil {
			fmt.Println("Digite um número inteiro, por favor.")
			continue
		}
		return num
	}
}

func isPrime(slice_primes []int, num int) bool {
	for _, value := range slice_primes {
		switch {
		case value*value > num:
			return true
		case num%value == 0:
			return false
		}
	}
	return true
}
