package tps

import (
	"math/big"
	"sync"
)

type SupplyEvent struct {
}

type BorrowEvent struct {
	BlockNumber *big.Int `json:"blockNumber"`
	Reserve     string   `json:"reserve"`
	OnBehalfOf  string   `json:"onBehalfOf"`
	User        string   `json:"user"`
	Amount      *big.Int `json:"amount"`
}

type WithdrawEvent struct {
	BlockNumber *big.Int `json:"blockNumber"`
	Reserve     string   `json:"reserve"`
	To          string   `json:"to"`
	User        string   `json:"user"`
	Amount      *big.Int `json:"amount"`
}

type RepayEvent struct {
}

type TransferATokenEvent struct {
	BlockNumber *big.Int `json:"blockNumber"`
	Reserve     string   `json:"reserve"`
	From        string   `json:"from"`
	To          string   `json:"to"`
	Amount      *big.Int `json:"amount"`
}

type ActiveUsersAggregator struct {
	mu       sync.Mutex
	activity map[int64][]string
}
