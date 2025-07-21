package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
}
