package main

import (
	"errors"
	"fmt"
)

// User defines our data structure
type User struct {
	Name string
	Age  int
}

// NewUser acts as our constructor function.
// It is idiomatic to return a pointer (*User) to avoid copying data in memory.
func NewUser(name string, age int) (*User, error) {
	// 1. Validation Logic
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if age < 0 {
		return nil, errors.New("age cannot be negative")
	}

	// 2. Encapsulated Initialization 
	u := User{
		Name: name,
		Age:  age,
	}

	// Go safely handles returning a pointer to local variables (escape analysis)
	return &u, nil
}

func main() {
	// Success case
	user1, err := NewUser("Alice", 30)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Created User: %s, Age: %d\n", user1.Name, user1.Age)
	}

	// Failure case
	_, err = NewUser("", -5)
	if err != nil {
		fmt.Println("Validation failed as expected:", err)
	}
}
