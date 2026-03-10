package types

import (
	"time"
)

// RawEventDBID represents a unique database ID for a raw event.
type RawEventDBID uint64

// RawEvent represents an uninterpreted EVM event log captured from a DAO contract.
type RawEvent struct {
	DBID              RawEventDBID
	DaoID             DAOID
	ChainID           BlockchainID
	ContractAddressID AddressID
	TxHash            string
	BlockHeight       uint64
	LogIndex          uint
	Timestamp         time.Time
	Topics            []string
	Data              string
}

// NewRawEvent creates a new RawEvent with the given IDs.
func NewRawEvent(daoID DAOID, chainID BlockchainID, contractAddressID AddressID) *RawEvent {
	return &RawEvent{
		DaoID:             daoID,
		ChainID:           chainID,
		ContractAddressID: contractAddressID,
	}
}
