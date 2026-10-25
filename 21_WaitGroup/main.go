package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		// wg.Go automatically calls wg.Add(1) and schedules defer wg.Done()
		wg.Go(func() {
			fmt.Printf("Worker %d starting\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done\n", i)
		})
	}

	// Blocks until all 3 workers are complete
	wg.Wait() 
	fmt.Println("All workers finished.")
}
