package tps

import (
	"math/big"
	"sync"
	"time"
)

type HfDropBlock struct {
	User             string    `json:"user"`
	LiquidationBlock *big.Int  `json:"liquidationBlock"`
	LowerBound       *big.Int  `json:"lowerBound"`
	UpperBound       *big.Int  `json:"upperBound"`
	DropBlock        *big.Int  `json:"dropBlock"`
	DropBlockTime    time.Time `json:"dropBlockTime"`
	UserCategory     int       `json:"userCategory"`
	Error            string    `json:"error"`
}

type HfDropBlockAggregator struct {
	mu      sync.Mutex
	Records []HfDropBlock
}

func (agg *HfDropBlockAggregator) Extend(rec HfDropBlock) {
	agg.mu.Lock()
	defer agg.mu.Unlock()
	agg.Records = append(agg.Records, rec)
}
