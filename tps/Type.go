package tps

import (
	"errors"
	"math/big"
	"slices"
	"sync"
	"time"
)

type LiquidationRecord struct {
	Day                        time.Time `json:"day"`
	BlockNumber                *big.Int  `json:"blockNumber"`
	CollateralAsset            string    `json:"collateralAsset"`
	DebtAsset                  string    `json:"debtAsset"`
	User                       string    `json:"user"`
	DebtToCover                *big.Int  `json:"debtToCover"`
	LiquidatedCollateralAmount *big.Int  `json:"liquidatedCollateralAmount"`
	Liquidator                 string    `json:"liquidator"`
}

// Returns true if block number of hfrec is lower than block number of r, else false
func (liqrec LiquidationRecord) IsLess(r LiquidationRecord) bool {
	return liqrec.BlockNumber.Cmp(r.BlockNumber) < 0
}

// Returns the record with the lowest block among a slice of records
func min(r []LiquidationRecord) LiquidationRecord {
	min := r[0]
	for _, rec := range r {
		if rec.IsLess(min) {
			min = rec
		}
	}
	return min
}

// Returns slice of unique users inside slice of liquidation records
func getusers(r []LiquidationRecord) []string {
	var allusr []string
	for _, rec := range r {
		if !slices.Contains(allusr, rec.User) {
			allusr = append(allusr, rec.User)
		}
	}
	return allusr
}

// Filter a liquidation records slice that correspond to specific user
func filter(r []LiquidationRecord, usr string) []LiquidationRecord {
	var fr []LiquidationRecord
	for _, rec := range r {
		if rec.User == usr {
			fr = append(fr, rec)
		}
	}
	return fr
}

// Return first liquidation for each user in r
func GetFirstLiqPerUser(r []LiquidationRecord) []LiquidationRecord {
	var uniqueliq []LiquidationRecord
	users := getusers(r)
	for _, usr := range users {
		usrliqrec := min(filter(r, usr))
		uniqueliq = append(uniqueliq, usrliqrec)
	}
	return uniqueliq
}

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

func (rec *UserHfRecord) IsLess(r *UserHfRecord) bool {
	return rec.BlockNumber.Cmp(r.BlockNumber) < 0
}

func (rec *UserHfRecord) IsHealthy() bool {
	return rec.HealthFactor.Cmp(big.NewInt(1000000000000000000)) > 0
}

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

// Slice of records with safe access
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

type HfDropBlock struct {
	User             string   `json:"user"`
	LiquidationBlock *big.Int `json:"liquidationBlock"`
	LowerBound       *big.Int `json:"lowerBound"`
	UpperBound       *big.Int `json:"upperBound"`
	DropBlock        *big.Int `json:"dropBlock"`
	Error            string   `json:"error"`
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

type PriceReturnRecord struct {
	BlockNumber     *big.Int `json:"blockNumber"`
	UnderlyingToken string   `json:"underlyingToken"`
	Price           *big.Int `json:"price"`
	NextBlockPrice  *big.Int `json:"nextBlockPrice"`
}
