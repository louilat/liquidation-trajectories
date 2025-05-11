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

type ActivityRecord struct {
	User        string
	BlockNumber *big.Int
}

func Contains(activity []ActivityRecord, user string, block *big.Int) bool {
	for _, a := range activity {
		if a.User == user && a.BlockNumber == block {
			return true
		}
	}
	return false
}

// func (a ActivityRecord) IsLess(b ActivityRecord) bool {
// 	return a.BlockNumber.Cmp(b.BlockNumber) < 0
// }

func (a ActivityRecord) IsLessThanBlock(b *big.Int) bool {
	return a.BlockNumber.Cmp(b) < 0
}

type ActiveUsersAggregator struct {
	mu       sync.Mutex
	Activity map[int64][]ActivityRecord
}

func (agg *ActiveUsersAggregator) Extend(k int64, v []ActivityRecord) {
	agg.mu.Lock()
	defer agg.mu.Unlock()
	agg.Activity[k] = v
}
