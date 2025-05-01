package returns

import (
	"liquidation-trajectories/tps"
	"math/big"
	"reflect"
	"testing"
)

func TestGetHfDropBounds(t *testing.T) {
	input := []tps.UserHfRecord{
		{
			User:             "0xb90594ea5128a8178e132286DC2B7fBaC7d7266c",
			LiquidationBlock: big.NewInt(22373250),
			BlockNumber:      big.NewInt(22373100),
			HealthFactor:     big.NewInt(1200000000000000000),
		},
		{
			User:             "0xb90594ea5128a8178e132286DC2B7fBaC7d7266c",
			LiquidationBlock: big.NewInt(22373250),
			BlockNumber:      big.NewInt(22373150),
			HealthFactor:     big.NewInt(1100000000000000000),
		},
		{
			User:             "0xb90594ea5128a8178e132286DC2B7fBaC7d7266c",
			LiquidationBlock: big.NewInt(22373250),
			BlockNumber:      big.NewInt(22373200),
			HealthFactor:     big.NewInt(1050000000000000000),
		},
		{
			User:             "0xb90594ea5128a8178e132286DC2B7fBaC7d7266c",
			LiquidationBlock: big.NewInt(22373250),
			BlockNumber:      big.NewInt(22373250),
			HealthFactor:     big.NewInt(1200000000000000000),
		},
	}

	actualOutput, err := getHfDropBounds(input)
	if err != nil {
		t.Errorf("received an undexpected error: %v", err)
	}

	expectedOutput := tps.HfDropBlock{
		User:             "0xb90594ea5128a8178e132286DC2B7fBaC7d7266c",
		LiquidationBlock: big.NewInt(22373250),
		LowerBound:       big.NewInt(22373200),
		UpperBound:       big.NewInt(22373250),
	}

	if !reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Errorf("expected output to be %v, but got %v", expectedOutput, actualOutput)
	}
}
