package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (client *Client) BroadcastTransaction(request TransactionRequest) (string, error) {
	url := "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"
	
	payload := TransactionRequest{
		Symbol:    request.Symbol,
		Price:     request.Price,
		Timestamp: uint64(time.Now().Unix()),
	}
	
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("cannot serialize payload: %v", err)
	}
	
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("cannot POST broadcast transaction: %v", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read response: %v", err)
	}
	
	var transactionResponse TransactionResponse
	err = json.Unmarshal(body, &transactionResponse)
	if err != nil {
		return "", fmt.Errorf("cannot deserialize response: %v", err)
	}
	
	return transactionResponse.TxHash, nil
}