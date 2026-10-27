package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()         // 1. Acquire exclusive access
	defer c.mu.Unlock() // 2. Ensure it unlocks when the function exits
	c.value++           // 3. Mutate the shared resource
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := SafeCounter{}

	// Spawn 1000 goroutines to safely increment the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("Final Count:", counter.Value()) // Always outputs 1000
}
