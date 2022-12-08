package types

import "time"

type Transaction struct {
	Sender   string  `json:"sender"`
	Recipent string  `json:"recipent"`
	Amount   float64 `json:"amount"`
}

type Block struct {
	Index        int           `json:"index"`
	Timestamp    time.Time     `json:"timestamp"`
	Transactions []Transaction `json:"transaction"`
	Proof        int           `json:"proof"`
	PreviousHash string        `json:"previous_hash"`
}
