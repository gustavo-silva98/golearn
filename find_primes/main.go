package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	fmt.Println(args[1])
	num, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(args)
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
	fmt.Println(primes)
	fmt.Printf("Tempo percorrido: %v\n", elapsed)

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
