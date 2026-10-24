package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

// composition: customer fields are promoted to order
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time 
	customer
}

func main() {
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
	}

	// Because of field promotion, you can also write this as: newOrder.name = "robin"
	newOrder.customer.name = "robin"
	
	fmt.Println(newOrder)
}
