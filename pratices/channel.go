package pratices

import (
	"fmt"
	"sync"
)

// Worker pool pattern for processing jobs efficiently.
// Buffered vs. Unbuffered Channels – know when to use which.
// Key concept: Avoid goroutine leaks by properly closing channels.
func RunChannel() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println("Result:", res)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2
	}
}
