package transaction

import "net/http"

type TransactionRequest struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type TransactionResponse struct {
	TxHash string `json:"tx_hash"`
}

type Client struct {
	httpClient *http.Client
}

// NewClient initializes a new Client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}