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
	// Começa o teste de channels e goroutines
	start := time.Now()
	results := get_pokemon_pages(6, 3)

	fmt.Println("\n--- Resultados Finais ---")
	var allCards []Card
	for body := range results { // Loop que lê do canal 'results' até ele ser fechado
		var resp APIResponse
		if err := json.Unmarshal(body, &resp); err != nil {
			log.Fatalf("Erro no Unmarshal: %v\n", err)
		}
		allCards = append(allCards, resp.Data...)
	}

	fmt.Printf("Total de cards carregados: %d\n", len(allCards))
	for idx, value := range allCards {
		fmt.Printf("%v| ID: %v - Name: %v - Types: %v - Subtypes: %v - HP: %v\n", idx+1, value.Id, value.Name, value.Types, value.Subtypes, value.HP)
	}

	fmt.Println("\nTodas as páginas foram processadas!")
	elapsed := time.Since(start)
	fmt.Printf("Tempo percorrido: %v\n", elapsed)
}

func worker(id int, jobs <-chan int, results chan<- []byte, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Trabalhador %d: Iniciando página %d\n", id, job)
		url := fmt.Sprintf("https://api.pokemontcg.io/v2/cards?page=%v&pageSize=250", job)
		byteReq := get_http(url)
		fmt.Printf("Trabalhador %d: Finalizou a página %d\n", id, job)

		results <- byteReq
	}
}

func get_http(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição na %v : Err = %v\n", url, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler body da url %v : Err = %v\n", url, err)
	}
	return body

}

func get_pokemon_pages(numPages int, numWorkers int) chan []byte {

	// Criação de channels
	jobs := make(chan int, numPages)
	results := make(chan []byte, numPages)

	var wg sync.WaitGroup

	// Criação de workers e lançamento de Goroutines
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Lançamento dos indices das páginas no channel Jovs
	for p := 1; p <= numPages; p++ {
		jobs <- p
	}
	close(jobs) // Fecha canal depois de inserir todos os jobs

	wg.Wait() // Trava a execução p/ syncronizar todos os workers

	close(results) // Fecha o channel de results depois de finalizar os workers

	return results
}
