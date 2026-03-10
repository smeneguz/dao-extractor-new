package rpc

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/dao-portal/flux/node"
	"github.com/dao-portal/flux/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/rs/zerolog"

	evmtypes "github.com/dao-portal/extractor/flux/evm/types"
	"github.com/dao-portal/extractor/utils"
)

var _ node.Node = &Node{}

// Node represents a node instance capable of fetching data from an EVM node
// trough JSON-RPC calls.
type Node struct {
	cfg       Config
	logger    zerolog.Logger
	chainID   string
	ethClient *ethclient.Client
}

func NewNode(ctx context.Context, logger zerolog.Logger, cfg Config) (*Node, error) {
	rpcClient, err := ethrpc.DialOptions(ctx, cfg.URL, ethrpc.WithHTTPClient(&http.Client{
		Timeout: cfg.RequestTimeout,
	}))
	if err != nil {
		return nil, fmt.Errorf("create rpc client: %w", err)
	}
	client := ethclient.NewClient(rpcClient)

	// Fetch the chain ID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id")
	}
	chainIDHex := fmt.Sprintf("0x%x", chainID)

	return &Node{
		cfg:       cfg,
		logger:    logger.With().Str(NodeType, cfg.URL).Logger(),
		ethClient: client,
		chainID:   chainIDHex,
	}, nil
}

// GetBlock implements node.Node.
func (n *Node) GetBlock(ctx context.Context, height types.Height) (types.Block, error) {
	// Get the block header
	ethBlock, err := n.ethClient.HeaderByNumber(ctx, height.ToBigInt())
	if err != nil {
		return nil, fmt.Errorf("get eth block header: %w", err)
	}

	// Get the block receipts
	blockNumber := ethrpc.BlockNumber(int64(height))
	receipts, err := n.ethClient.BlockReceipts(ctx, ethrpc.BlockNumberOrHash{
		BlockNumber: &blockNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("get block receipts: %w", err)
	}

	// Create the Tx objects from the block receipts
	txs := make([]*evmtypes.Tx, len(receipts))
	var logs evmtypes.Logs
	for i, r := range receipts {
		// Convert the logs to our format
		txLogs := utils.Map(r.Logs, LogFromEthlog)
		// Flatten the logs to have them all also in the block
		logs = append(logs, txLogs...)

		// Create the transaction instance
		txs[i] = TxFromReceipt(r).
			WithLogs(txLogs)
	}

	// Create the block
	return evmtypes.NewBlock(
		n.chainID,
		height,
		time.Unix(int64(ethBlock.Time), 0),
		txs,
	).WithLogs(logs), nil
}

// GetChainID implements node.Node.
func (n *Node) GetChainID() string {
	return n.chainID
}

// GetCurrentHeight implements node.Node.
func (n *Node) GetCurrentHeight(ctx context.Context) (types.Height, error) {
	height, err := n.ethClient.BlockNumber(ctx)
	return types.Height(height), err
}

// GetLowestHeight implements node.Node.
func (n *Node) GetLowestHeight(ctx context.Context) (types.Height, error) {
	// Try to get the first block
	_, err := n.ethClient.HeaderByNumber(ctx, big.NewInt(0))
	if err == nil {
		return 0, nil
	}

	// Get the current height
	currentHeight, err := n.GetCurrentHeight(ctx)
	if err != nil {
		return 0, err
	}

	// Start a binary search for the lowest available block
	low := types.Height(0)
	high := currentHeight
	lowestAvailable := currentHeight

	for low <= high {
		mid := (low + high) / 2
		if _, err := n.ethClient.HeaderByNumber(ctx, mid.ToBigInt()); err == nil {
			lowestAvailable = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return lowestAvailable, nil
}

// GetEthClient get the go-ethereum client instance.
func (n *Node) GetEthClient() *ethclient.Client {
	return n.ethClient
}

// GetTxGasFeesAndUsed get the gas fees and gas used for a transaction.
func (n *Node) GetTxGasFeesAndUsed(ctx context.Context, tx *evmtypes.Tx) (*big.Int, *big.Int, error) {
	// Get the transaction receipt to get the transaction fees and gas used
	client := n.GetEthClient()
	transaction, err := client.TransactionReceipt(ctx, common.BytesToHash(tx.Hash))
	if err != nil {
		return nil, nil, fmt.Errorf("get transaction receipt: %w", err)
	}

	// Get the gas used and fees for the transaction
	gasUsed := new(big.Int).SetUint64(transaction.GasUsed)
	gasFees := new(big.Int).Mul(transaction.EffectiveGasPrice, gasUsed)

	// Get the gas used and fees for the blob transaction
	blobGasUsed := new(big.Int).SetUint64(transaction.BlobGasUsed)
	blobGasFees := big.NewInt(0)
	if transaction.BlobGasPrice != nil {
		blobGasFees = new(big.Int).Mul(transaction.BlobGasPrice, blobGasUsed)
	}

	// Total gas used and fees
	gasFees = new(big.Int).Add(gasFees, blobGasFees)
	gasUsed = new(big.Int).Add(gasUsed, blobGasUsed)

	return gasFees, gasUsed, nil
}
