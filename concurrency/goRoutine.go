package concurrency

import (
	"fmt"
	"sync"
)

// Goroutines are lightweight, but excessive spawning can lead to high memory usage.
// Use sync.WaitGroup to manage goroutines efficiently.
// Key concept: Avoid common pitfalls like capturing loop variables incorrectly in goroutines.
func GoRoutine() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) { //<- Capture i correctly
			defer wg.Done()
			fmt.Println("Goroutine", i)
		}(i) // Pass i as argument
	}
	wg.Wait()

}
