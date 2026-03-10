package tokenevents

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Event signature hashes for ERC20Votes token events.
var (
	// Transfer(address indexed from, address indexed to, uint256 value)
	transferEventSig = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

	// DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
	delegateChangedEventSig = crypto.Keccak256Hash([]byte("DelegateChanged(address,address,address)"))

	// DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
	delegateVotesChangedEventSig = crypto.Keccak256Hash([]byte("DelegateVotesChanged(address,uint256,uint256)"))
)

const (
	eventNameTransfer             = "Transfer"
	eventNameDelegateChanged      = "DelegateChanged"
	eventNameDelegateVotesChanged = "DelegateVotesChanged"
)

// identifyEvent returns the event name based on topic0, or empty string if unknown.
func identifyEvent(topic0 common.Hash) string {
	switch topic0 {
	case transferEventSig:
		return eventNameTransfer
	case delegateChangedEventSig:
		return eventNameDelegateChanged
	case delegateVotesChangedEventSig:
		return eventNameDelegateVotesChanged
	default:
		return ""
	}
}

// parseTransferEvent extracts from, to, value from an ERC20 Transfer log.
// Transfer has 2 indexed params (from, to) and 1 non-indexed param (value).
// Topics: [sig, from, to], Data: [value]
func parseTransferEvent(topics []common.Hash, data []byte) (from, to common.Address, value *big.Int, ok bool) {
	if len(topics) < 3 || len(data) < 32 {
		return common.Address{}, common.Address{}, nil, false
	}
	from = common.BytesToAddress(topics[1].Bytes())
	to = common.BytesToAddress(topics[2].Bytes())
	value = new(big.Int).SetBytes(data[:32])
	return from, to, value, true
}

// parseDelegateChangedEvent extracts delegator, fromDelegate, toDelegate.
// All 3 params are indexed.
// Topics: [sig, delegator, fromDelegate, toDelegate], Data: []
func parseDelegateChangedEvent(topics []common.Hash) (delegator, fromDelegate, toDelegate common.Address, ok bool) {
	if len(topics) < 4 {
		return common.Address{}, common.Address{}, common.Address{}, false
	}
	delegator = common.BytesToAddress(topics[1].Bytes())
	fromDelegate = common.BytesToAddress(topics[2].Bytes())
	toDelegate = common.BytesToAddress(topics[3].Bytes())
	return delegator, fromDelegate, toDelegate, true
}

// parseDelegateVotesChangedEvent extracts delegate, previousBalance, newBalance.
// 1 indexed param (delegate), 2 non-indexed params (previousBalance, newBalance).
// Topics: [sig, delegate], Data: [previousBalance, newBalance]
func parseDelegateVotesChangedEvent(topics []common.Hash, data []byte) (delegate common.Address, previousBalance, newBalance *big.Int, ok bool) {
	if len(topics) < 2 || len(data) < 64 {
		return common.Address{}, nil, nil, false
	}
	delegate = common.BytesToAddress(topics[1].Bytes())
	previousBalance = new(big.Int).SetBytes(data[:32])
	newBalance = new(big.Int).SetBytes(data[32:64])
	return delegate, previousBalance, newBalance, true
}
