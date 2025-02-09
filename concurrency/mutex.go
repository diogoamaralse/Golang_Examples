package concurrency

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex // Mutex to protect the shared counter
)

func Mutex() {
	var wg sync.WaitGroup

	// Launch 10 goroutines to increment the counter concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the result after all increments
	fmt.Println("Final Counter:", counter)
}

func increment() {
	mu.Lock()         // Lock the critical section
	defer mu.Unlock() // Ensure the lock is released after function execution
	counter++
}
