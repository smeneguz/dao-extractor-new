package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

// EVMNode defines the interface for EVM-specific node operations
// that governor modules need beyond the basic flux Node interface.
type EVMNode interface {
	GetEthClient() *ethclient.Client
	GetTxGasFeesAndUsed(ctx context.Context, tx *Tx) (*big.Int, *big.Int, error)
}
