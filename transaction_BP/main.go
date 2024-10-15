package main

import (
	"fmt"
	"time"
	"transaction_BP/transaction"
)

func main() {
	client := transaction.NewClient()

	// Transaction Data Sample
	reqBody := transaction.TransactionRequest{
		Symbol: "BTC",
		Price:  50000,
	}

	// Broadcast the Transaction
	txHash, err := client.BroadcastTransaction(reqBody)
	if err != nil {
		fmt.Printf("Cannot broadcasting transaction: %v\n", err)
		return
	}

	fmt.Printf("Transaction broadcasted with hash: %s\n", txHash)

	// Monitor the Transaction Status
	status, err := client.MonitorTransaction(txHash, 5*time.Second)
	if err != nil {
		fmt.Printf("Cannot monitoring transaction: %v\n", err)
		return
	}

	fmt.Printf("Transaction finalized with status: %s\n", status)
}
