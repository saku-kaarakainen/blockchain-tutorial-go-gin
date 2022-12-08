package blockchain

import (
	"blockchain-tutorial-go-gin/types"
	"time"
)

var blockchain []types.Block
var currentTransactions []types.Transaction

func CreateGenesis() types.Block {
	return NewBlock(1, "1")
}

func NewBlock(proof int, previoushHash string) types.Block {
	var timestamp = time.Now()
	block := types.Block{
		Index:        len(blockchain) + 1,
		Timestamp:    timestamp,
		Transactions: currentTransactions,
		Proof:        proof,
		PreviousHash: previoushHash,
	}

	blockchain = append(blockchain, block)
	currentTransactions = nil

	return block
}

func NewTransaction(
	sender string,
	recipient string,
	amount float64) int {

	newTransaction := types.Transaction{
		Sender:   sender,
		Recipent: recipient,
		Amount:   amount,
	}

	append(currentTransactions, newTransaction)

	return len(blockchain) + 1
}
