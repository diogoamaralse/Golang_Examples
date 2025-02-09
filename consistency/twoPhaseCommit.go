package consistency

import (
	"fmt"
	"time"
)

//Two-Phase Commit (2PC) Example for Distributed Transactions
//The two-phase commit (2PC) protocol works by having the transaction coordinator and participant services agree on whether to commit or abort a transaction. The first phase is a vote phase, and the second phase is a commit phase.
//Steps in Two-Phase Commit:
//Phase 1: The coordinator sends a prepare request to the participants (services).
//Phase 2: If all participants vote commit, the coordinator sends a commit request. If any participant votes abort, the coordinator sends an abort request.

func TwoPhaseCommit() {
	// Simulating distributed services
	accountService := &AccountService{balance: 1000}
	receiptService := &ReceiptService{receiptID: 1}

	// Create a Payment Service that communicates with both services
	paymentService := &PaymentService{
		accountService: accountService,
		receiptService: receiptService,
	}

	// Attempt a payment transaction
	fmt.Println("Starting transaction...")

	// Example payment amount
	amount := 500

	// Process the transaction
	paymentService.processTransaction(amount)

	// Final state of the account after transaction
	fmt.Println("Final account balance:", accountService.balance)
}

// Simulated distributed system services
type AccountService struct {
	balance int
}

type ReceiptService struct {
	receiptID int
}

type PaymentService struct {
	accountService *AccountService
	receiptService *ReceiptService
}

// Prepare phase: Checks if AccountService and ReceiptService are ready to commit the transaction
func (s *AccountService) preparePayment(amount int) bool {
	// Simulate account balance check before transaction
	if s.balance >= amount {
		fmt.Println("Account service: Balance sufficient for transaction.")
		return true
	}
	fmt.Println("Account service: Insufficient funds.")
	return false
}

func (s *ReceiptService) prepareReceipt(transactionID int) bool {
	// Simulate preparing a receipt
	fmt.Println("Receipt service: Preparing receipt for transaction", transactionID)
	return true
}

// 2PC Phase 1: Prepare Phase
func (ps *PaymentService) prepareTransaction(amount int) bool {
	// Phase 1: Ask each service if they are ready to commit
	if !ps.accountService.preparePayment(amount) {
		return false
	}

	if !ps.receiptService.prepareReceipt(amount) {
		return false
	}

	return true
}

// 2PC Phase 2: Commit or Abort Phase
func (ps *PaymentService) commitTransaction(amount int) bool {
	// Phase 2: Commit the transaction if all services are ready
	ps.accountService.balance -= amount
	fmt.Println("Payment service: Transaction committed. Account balance is now", ps.accountService.balance)
	return true
}

// 2PC: Handle transaction with retries for consistency
func (ps *PaymentService) processTransaction(amount int) {
	// Phase 1: Prepare transaction
	fmt.Println("Payment service: Starting Phase 1 - Prepare transaction.")
	if !ps.prepareTransaction(amount) {
		fmt.Println("Payment service: Transaction failed in Phase 1. Aborting.")
		return
	}

	// Simulate a potential failure in the network or service during commit phase
	time.Sleep(1 * time.Second) // Simulate delay

	// Phase 2: Commit transaction
	fmt.Println("Payment service: Starting Phase 2 - Commit transaction.")
	if ps.commitTransaction(amount) {
		fmt.Println("Payment service: Transaction completed successfully.")
	} else {
		fmt.Println("Payment service: Transaction failed in Phase 2. Rolling back.")
	}
}
