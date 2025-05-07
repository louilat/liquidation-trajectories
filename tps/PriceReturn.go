package tps

import "math/big"

type PriceReturnRecord struct {
	BlockNumber     *big.Int `json:"blockNumber"`
	UnderlyingToken string   `json:"underlyingToken"`
	Price           *big.Int `json:"price"`
	NextBlockPrice  *big.Int `json:"nextBlockPrice"`
}
