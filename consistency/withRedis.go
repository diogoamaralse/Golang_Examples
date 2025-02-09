package consistency

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

const campaignBudgetLimit = 10000

var currentSpend int

func ProcessWithRedis() {
	// Initialize Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Default DB
	})

	// Initialize campaign state
	rdb.Set(ctx, "campaign_total_spend", 0, 0)   // Initial total spend is 0
	rdb.Set(ctx, "campaign_status", "active", 0) // Campaign is active initially

	// Simulating cashback requests
	var wg sync.WaitGroup
	wg.Add(5)

	users := []struct {
		userID         string
		cashbackAmount int
	}{
		{"user1", 1000},
		{"user2", 1500},
		{"user3", 2000},
		{"user4", 3000},
		{"user5", 1500},
	}

	for _, user := range users {
		go processCashback(user.userID, user.cashbackAmount, &wg)
	}

	wg.Wait()

	// Final state of the campaign
	finalTotalSpend, err := rdb.Get(ctx, "campaign_total_spend").Int()
	if err != nil {
		log.Printf("Error fetching final total spend: %v", err)
	}
	fmt.Printf("Final total spend: €%d\n", finalTotalSpend)

	// Check if the campaign is finished
	status, err := rdb.Get(ctx, "campaign_status").Result()
	if err != nil {
		log.Printf("Error fetching campaign status: %v", err)
	}
	fmt.Println("Campaign status:", status)
}

// Simulated service processing cashback requests
func processCashback(userID string, cashbackAmount int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire a lock before updating campaign state (to ensure no race condition)
	lockKey := "campaign_lock"
	ok, err := rdb.SetNX(ctx, lockKey, 1, 10*time.Second).Result()
	if err != nil {
		log.Printf("Error acquiring lock: %v", err)
		return
	}

	if !ok {
		log.Println("Lock not acquired, skipping...")
		return
	}
	defer rdb.Del(ctx, lockKey) // Ensure the lock is released after processing

	// Check if the remaining budget allows the cashback
	currentTotalSpend, err := rdb.Get(ctx, "campaign_total_spend").Int()
	if err != nil && err != redis.Nil {
		log.Printf("Error fetching current total spend: %v", err)
		return
	}

	// If the remaining budget is insufficient, cancel the cashback
	if currentTotalSpend+cashbackAmount > campaignBudgetLimit {
		log.Printf("Insufficient funds. Cannot process cashback for user %s.\n", userID)
		return
	}

	// Apply the cashback by updating the total spend
	_, err = rdb.IncrBy(ctx, "campaign_total_spend", int64(cashbackAmount)).Result()
	if err != nil {
		log.Printf("Error updating total spend: %v", err)
		return
	}

	fmt.Printf("Processed cashback of €%d for user %s. Current total spend: €%d\n", cashbackAmount, userID, currentTotalSpend+cashbackAmount)

	// If the total spend has reached or exceeded the limit, mark the campaign as finished
	if currentTotalSpend+cashbackAmount >= campaignBudgetLimit {
		fmt.Println("Campaign budget limit reached. Ending campaign.")
		rdb.Set(ctx, "campaign_status", "finished", 0) // Mark campaign as finished
	}
}
