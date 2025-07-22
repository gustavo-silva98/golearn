package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	Data []Card `json:"data"`
}

type Card struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Types    []string `json:"types"`
	Subtypes []string `json:"subtypes"`
	HP       string   `json:"hp"`
}

func main() {
	resp, err := http.Get("https://api.pokemontcg.io/v2/cards?page=1&pageSize=50")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var apiResp APIResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Fatal(err)
	}
	for idx, value := range apiResp.Data {
		fmt.Printf("%v| ID: %v - Name: %v - Types: %v - Subtypes: %v - HP: %v\n", idx+1, value.Id, value.Name, value.Types, value.Subtypes, value.HP)
	}
	// Começa o teste de channels e goroutines

	const numWorkers = 3
	const numPages = 10

	jobs := make(chan int, numPages)
	results := make(chan string, numPages)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for p := 1; p <= numPages; p++ {
		jobs <- p
	}
	close(jobs)

	wg.Wait()

	close(results)

	fmt.Println("\n--- Resultados Finais ---")
	for res := range results { // Loop que lê do canal 'results' até ele ser fechado
		fmt.Println(res)
	}

	fmt.Println("\nTodas as páginas foram processadas!")
}

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Trabalhador %d: Iniciando página %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Trabalhador %d: Finalizou a página %d\n", id, job)

		results <- fmt.Sprintf("Página %d processada pelo trabalhador %d", job, id)
	}
}
