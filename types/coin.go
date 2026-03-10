package types

import (
	"fmt"
	"math/big"
)

// Coin represents a coin that has been used in a transaction.
type Coin struct {
	// Denom is the denomination of the coin, in case of a non-native token
	// this will be the contract address.
	Denom string
	// Amount is the amount of the coin.
	Amount *big.Int
}

// NewCoin creates a new Coin instance.
func NewCoin(denom string, amount *big.Int) Coin {
	return Coin{
		Amount: amount,
		Denom:  denom,
	}
}

func NewCoinFromString(denom string, amount string) (Coin, error) {
	amountInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return Coin{}, fmt.Errorf("invalid big.Int amount: %s", amount)
	}
	return NewCoin(denom, amountInt), nil
}
