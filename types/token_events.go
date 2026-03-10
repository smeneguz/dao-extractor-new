package types

import (
	"math/big"
	"time"
)

// TokenTransfer represents an ERC20 Transfer event from a governance token contract.
type TokenTransfer struct {
	DBID         uint64
	DaoID        DAOID
	ChainID      string
	TokenAddress string
	FromAddress  string
	ToAddress    string
	Amount       *big.Int
	TxHash       string
	BlockHeight  uint64
	LogIndex     uint
	Timestamp    time.Time
}

// DelegationEvent represents a DelegateChanged event from a governance token contract.
type DelegationEvent struct {
	DBID         uint64
	DaoID        DAOID
	ChainID      string
	TokenAddress string
	Delegator    string
	FromDelegate string
	ToDelegate   string
	TxHash       string
	BlockHeight  uint64
	LogIndex     uint
	Timestamp    time.Time
}

// DelegateVotesChanged represents a DelegateVotesChanged event from a governance token contract.
type DelegateVotesChanged struct {
	DBID            uint64
	DaoID           DAOID
	ChainID         string
	TokenAddress    string
	Delegate        string
	PreviousBalance *big.Int
	NewBalance      *big.Int
	TxHash          string
	BlockHeight     uint64
	LogIndex        uint
	Timestamp       time.Time
}
