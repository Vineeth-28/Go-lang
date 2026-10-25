package main

import (
    "fmt"
    "time"
)

func computeTask(id int) {
    fmt.Printf("Task %d starting...\n", id)
    time.Sleep(500 * time.Millisecond) // Simulating work
    fmt.Printf("Task %d finished!\n", id)
}

func main(){
	//launching a named function as go routine 

	go computeTask(1);
	//launching an anonymous function as go routine 

	    go func(msg string) {
        fmt.Println(msg)
    }("Hello from anonymous goroutine!")

    // Giving goroutines time to print before main exits
    time.Sleep(1 * time.Second) }