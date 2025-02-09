package resilience

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sony/gobreaker"
)

//A Circuit Breaker is a useful pattern for preventing cascading failures in a distributed system, especially when services rely on external systems or APIs.
//It helps to "break" the connection if the failure threshold is met and prevents the system from trying to call a failing service repeatedly, thus reducing the load on the failing service and allowing it to recover.

func GoCircuitBreake() {
	// Configure the circuit breaker
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "ExampleCircuitBreaker",
		MaxRequests: 5,                // Maximum number of requests allowed before the circuit is checked
		Interval:    60 * time.Second, // How long the circuit breaker will stay open
		Timeout:     10 * time.Second, // Time to wait for the external service before considering it failed
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 2 // Break the circuit after 2 consecutive failures
		},
	})

	// Simulate multiple calls to the unreliable API
	for i := 0; i < 10; i++ {
		fmt.Printf("Attempt #%d...\n", i+1)

		// Use the circuit breaker to make the API call
		result, err := cb.Execute(func() (interface{}, error) {
			err := unreliableAPI()
			if err != nil {
				fmt.Println("API call failed:", err)
				return nil, err // Return `nil` for the result and the error
			}
			fmt.Println("API call succeeded")
			return "Success", nil // Return a valid result and `nil` for the error
		})

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			// This part will be triggered if the function executes successfully.
			fmt.Println("Result:", result)
		}

		time.Sleep(1 * time.Second) // Simulate time between requests
	}
}

// Simulated API call that may fail
func unreliableAPI() error {
	// Randomly simulate failure or success
	if rand.Float32() < 0.5 {
		return fmt.Errorf("API call failed")
	}
	return nil
}
