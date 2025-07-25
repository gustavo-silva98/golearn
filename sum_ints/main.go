package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Order struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Item    string  `json:"item"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, event json.RawMessage) (int, error) {
	var order Order
	if err := json.Unmarshal(event, &order); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return 0, err
	}
	return sum_ints(5, 5), nil
}

func sum_ints(a int, b int) int {
	return a + b
}
