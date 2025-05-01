package tps

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFilterHealthyFactors(t *testing.T) {
	input := []UserHfRecord{
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
			HealthFactor:     big.NewInt(100000000000000000),
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

	actualOutput := FilterHealthyFactors(input)

	expectedOutput := []UserHfRecord{
		{
			User:             "0xb90594ea5128a8178e132286DC2B7fBaC7d7266c",
			LiquidationBlock: big.NewInt(22373250),
			BlockNumber:      big.NewInt(22373100),
			HealthFactor:     big.NewInt(1200000000000000000),
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

	if !reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Errorf("expected output to be %v, but got %v", expectedOutput, actualOutput)
	}
}

func TestArgMax(t *testing.T) {

	input := []UserHfRecord{
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
			HealthFactor:     big.NewInt(100000000000000000),
		},
		{
			User:             "0xb90594ea5128a8178e132286DC2B7fBaC7d7266c",
			LiquidationBlock: big.NewInt(22373250),
			BlockNumber:      big.NewInt(22373200),
			HealthFactor:     big.NewInt(1350000000000000000),
		},
		{
			User:             "0xb90594ea5128a8178e132286DC2B7fBaC7d7266c",
			LiquidationBlock: big.NewInt(22373250),
			BlockNumber:      big.NewInt(22373250),
			HealthFactor:     big.NewInt(1200000000000000000),
		},
		{
			User:             "0xb90594ea5128a8178e132286DC2B7fBaC7d7266c",
			LiquidationBlock: big.NewInt(22373250),
			BlockNumber:      big.NewInt(22373230),
			HealthFactor:     big.NewInt(1200000000000000000),
		},
	}

	actualOutput, err := ArgMax(input)
	if err != nil {
		t.Errorf("0 raised an unexpected error for actualOutput")
	}

	expectedOutput := 3

	if actualOutput != expectedOutput {
		t.Errorf("0 expected output to be %v, but got %v", expectedOutput, actualOutput)
	}

	input1 := []UserHfRecord{}

	_, err = ArgMax(input1)
	if err == nil {
		t.Errorf("1 was expecting an error, received nil")
	}

}
