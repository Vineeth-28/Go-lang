package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// os.ReadFile loads the entire file into a byte slice ([]byte)
	content, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	// Convert bytes to string to display it
	fmt.Println(string(content))
}
