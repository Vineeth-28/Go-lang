package main

import (
	"log"
	"os"
)

func main() {
	message := []byte("Hello, Go!")
	
	// 0644 sets Unix file permissions (Read/Write for owner, Read for others)
	err := os.WriteFile("output.txt", message, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}
}
