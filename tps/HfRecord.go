package tps

import (
	"errors"
	"math/big"
	"sync"
)

type UserHfRecord struct {
	User                        string   `json:"user"`
	LiquidationBlock            *big.Int `json:"liquidationBlock"`
	BlockNumber                 *big.Int `json:"blockNumber"`
	Timestamp                   *big.Int `json:"timestamp"`
	TotalCollateralBase         *big.Int `json:"totalCollateralBase"`
	TotalDebtBase               *big.Int `json:"totalDebtBase"`
	AvailableBorrowsBase        *big.Int `json:"availableBorrowsBase"`
	CurrentLiquidationThreshold *big.Int `json:"currentLiquidationThreshold"`
	Ltv                         *big.Int `json:"ltv"`
	HealthFactor                *big.Int `json:"healthFactor"`
}

// Compare to UserHfRecord based on blockNumber
func (rec *UserHfRecord) IsLess(r *UserHfRecord) bool {
	return rec.BlockNumber.Cmp(r.BlockNumber) < 0
}

// Return true if hf > 1, else false
func (rec *UserHfRecord) IsHealthy() bool {
	return rec.HealthFactor.Cmp(big.NewInt(1000000000000000000)) > 0
}

// Return the index of last UserHfRecord with hf > 1
func ArgMax(r []UserHfRecord) (int, error) {
	if len(r) == 0 {
		return 0, errors.New("cannot find argmax of empty slice")
	}
	max := 0
	for i, rec := range r {
		if r[max].IsLess(&rec) && rec.IsHealthy() {
			max = i
		}
	}
	if max > 0 || r[0].IsHealthy() {
		return max, nil
	}
	return 0, errors.New("could not find healthy hf in slice")
}

// Map of records with safe access
type UserHfRecordAggregator struct {
	mu      sync.Mutex
	Records map[LiquidationRecord][]UserHfRecord
}

// Extends safely a slice of records
func (agg *UserHfRecordAggregator) Extend(key LiquidationRecord, val []UserHfRecord) {
	agg.mu.Lock()
	defer agg.mu.Unlock()
	agg.Records[key] = val
}
