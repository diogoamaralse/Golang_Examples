package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Worker pool pattern for processing jobs efficiently.
// Buffered vs. Unbuffered Channels – know when to use which.
// Key concept: Avoid goroutine leaks by properly closing channels.

//Some important notes:
//buffered channel as example make(chan int, 5)
//Storage: Holds up to 5 values temporarily
//Order: FIFO (First-In-First-Out)
//Concurrency: Safe for goroutines, blocks when full/empty
//Blocking Behavior: Sender blocks when full, receiver blocks when empty

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
	close(results) // Prevents workers from waiting indefinitely

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

func ChannelGoRoutine() {
	ch := make(chan string, 1) // Buffered channel with capacity 1

	go workerGoRoutine(ch) // Start goroutine

	fmt.Println("Waiting for worker...")
	msg := <-ch // Receive message (blocks if empty)
	fmt.Println(msg)
}
func workerGoRoutine(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Task done!" // Send message to channel
}
