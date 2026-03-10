package types

import (
	"database/sql/driver"
	"fmt"
	"math/big"
)

func deserializeBigInt(src any) (*big.Int, error) {
	switch src := src.(type) {
	case string:
		value, ok := new(big.Int).SetString(src, 10)
		if !ok {
			return nil, fmt.Errorf("invalid big.Int string: %s", src)
		}
		return value, nil
	case []byte:
		value, ok := new(big.Int).SetString(string(src), 10)
		if !ok {
			return nil, fmt.Errorf("invalid big.Int bytes: %s", src)
		}
		return value, nil
	case int64:
		return new(big.Int).SetInt64(src), nil
	default:
		return nil, fmt.Errorf("unsupported src type: %T", src)
	}
}

// BigNumeric represents a big.Int that can be used to store numeric data in a database.
type BigNumeric struct {
	Int *big.Int
}

func NewBigNumeric(i *big.Int) *BigNumeric {
	if i == nil {
		return nil
	}
	return &BigNumeric{
		Int: i,
	}
}

func (bn *BigNumeric) Value() (driver.Value, error) {
	if bn == nil {
		return nil, nil
	}
	return bn.Int.String(), nil
}

func (bn *BigNumeric) Scan(src any) error {
	value, err := deserializeBigInt(src)
	if err != nil {
		return err
	}
	bn.Int = value
	return nil
}

// -------------------------------------------------------------------------------------------------------------------
// ---- big.Int helper
// -------------------------------------------------------------------------------------------------------------------

type BigInt struct {
	*big.Int
}

func NewBigInt(i *big.Int) *BigInt {
	if i == nil {
		return nil
	}
	return &BigInt{
		Int: i,
	}
}

func (bn *BigInt) Value() (driver.Value, error) {
	if bn == nil {
		return nil, nil
	}
	return bn.Int.String(), nil
}

func (bn *BigInt) Scan(src any) error {
	value, err := deserializeBigInt(src)
	if err != nil {
		return err
	}
	bn.Int = value
	return nil
}
