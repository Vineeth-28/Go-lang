package main

import (
	"fmt"
	"time"
)

// Order represents the blueprint for a customer purchase.
type Order struct {
	ID        string
	Amount    float32
	Status    string
	CreatedAt time.Time
}

func main() {
	// Instance 1: Initializing and then setting a field later
	order1 := Order{
		ID:     "1",
		Amount: 500.0,
		Status: "Received",
	}
	order1.CreatedAt = time.Now()

	// Accessing a specific field value
	fmt.Println("Order 1 Status:", order1.Status)

	// Instance 2: Full inline initialization
	order2 := Order{
		ID:        "2",
		Amount:    750.5,
		Status:    "Processing",
		CreatedAt: time.Now(),
	}

	// Printing the entire struct instance
	fmt.Println("Order 2 details:", order2)
	fmt.Println("Order 1 details:", order1)
}
