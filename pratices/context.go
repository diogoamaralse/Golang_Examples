package pratices

import (
	"context"
	"fmt"
	"time"
)

// Why? Prevent goroutines from running indefinitely.
// Key concept: Always call cancel() to free up resources.
func Context() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go doWork(ctx)
	time.Sleep(3 * time.Second)
}

func doWork(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work done")
	case <-ctx.Done():
		fmt.Println("Work cancelled:", ctx.Err())
	}
}
