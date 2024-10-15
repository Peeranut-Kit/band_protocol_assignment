package transaction

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// CheckTransactionStatus checks the status of a transaction by its hash
func (client *Client) CheckTransactionStatus(txHash string) (string, error) {
	url := fmt.Sprintf("https://mock-node-wgqbnxruha-as.a.run.app/check/%s", txHash)
	
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("cannot check transaction status: %v", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read response: %v", err)
	}
	
	return string(body), nil
}

// MonitorTransaction periodically checks the transaction status until it is confirmed, failed, or times out
func (client *Client) MonitorTransaction(txHash string, interval time.Duration) (string, error) {
	for {
		statusString, err := client.CheckTransactionStatus(txHash)
		if err != nil {
			return "", fmt.Errorf("cannot monitor transaction status: %v", err)
		}

		statusStringSegment := strings.Split(statusString, "\"")
		status := statusStringSegment[len(statusStringSegment)-2]

		fmt.Printf("Current status: %s\n", status)

		if status == "CONFIRMED" || status == "FAILED" || status == "DNE" {
			fmt.Printf("Current status: %s\n", status)
			return status, nil
		}

		time.Sleep(interval)
	}
}
