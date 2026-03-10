package fallback

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/dao-portal/flux/node"
	"github.com/dao-portal/flux/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"

	rpcnode "github.com/dao-portal/extractor/flux/evm/node/rpc"
	evmtypes "github.com/dao-portal/extractor/flux/evm/types"
)

var _ node.Node = &Node{}
var _ evmtypes.EVMNode = &Node{}

const maxFallbackRetries = 5

// Node wraps a primary and fallback RPC node. When the primary fails to
// fetch a block, the fallback is tried with rate limiting and automatic
// retry on 429 errors.
type Node struct {
	primary  *rpcnode.Node
	fallback *rpcnode.Node
	logger   zerolog.Logger

	cooldown time.Duration

	// Throttle fallback calls: only one at a time + cooldown between calls.
	mu       sync.Mutex
	lastCall time.Time
}

func NewNode(
	primary *rpcnode.Node,
	fallback *rpcnode.Node,
	logger zerolog.Logger,
	cooldown time.Duration,
) *Node {
	return &Node{
		primary:  primary,
		fallback: fallback,
		logger:   logger,
		cooldown: cooldown,
	}
}

// GetChainID implements node.Node.
func (n *Node) GetChainID() string {
	return n.primary.GetChainID()
}

// GetCurrentHeight implements node.Node.
func (n *Node) GetCurrentHeight(ctx context.Context) (types.Height, error) {
	return n.primary.GetCurrentHeight(ctx)
}

// GetLowestHeight implements node.Node.
func (n *Node) GetLowestHeight(ctx context.Context) (types.Height, error) {
	return n.primary.GetLowestHeight(ctx)
}

// GetBlock implements node.Node.
// Strategy:
//  1. Try primary node (fast, no rate limit)
//  2. If primary fails, try fallback with rate limiting
//  3. If fallback returns 429, wait with exponential backoff and retry
//     up to maxFallbackRetries times — the block is NEVER dropped at this level
func (n *Node) GetBlock(ctx context.Context, height types.Height) (types.Block, error) {
	block, err := n.primary.GetBlock(ctx, height)
	if err == nil {
		return block, nil
	}

	n.logger.Debug().
		Err(err).
		Uint64("height", uint64(height)).
		Msg("primary node failed, trying fallback")

	// Retry loop for the fallback node with exponential backoff on 429.
	var lastErr error
	for attempt := 0; attempt < maxFallbackRetries; attempt++ {
		// Throttle: enforce cooldown between fallback calls.
		n.waitCooldown(ctx)

		block, lastErr = n.fallback.GetBlock(ctx, height)
		if lastErr == nil {
			if attempt > 0 {
				n.logger.Debug().
					Uint64("height", uint64(height)).
					Int("attempt", attempt+1).
					Msg("fallback succeeded after retry")
			}
			return block, nil
		}

		// If not a rate limit error, no point retrying.
		if !isRateLimitError(lastErr) {
			return nil, lastErr
		}

		// Exponential backoff: 3s, 6s, 12s, 24s, 48s
		backoff := time.Duration(3<<uint(attempt)) * time.Second
		n.logger.Warn().
			Uint64("height", uint64(height)).
			Int("attempt", attempt+1).
			Dur("backoff", backoff).
			Msg("fallback rate limited (429), backing off")

		select {
		case <-time.After(backoff):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	return nil, lastErr
}

// waitCooldown enforces a minimum interval between fallback calls.
func (n *Node) waitCooldown(ctx context.Context) {
	n.mu.Lock()
	elapsed := time.Since(n.lastCall)
	if elapsed < n.cooldown {
		wait := n.cooldown - elapsed
		n.mu.Unlock()

		select {
		case <-time.After(wait):
		case <-ctx.Done():
			return
		}

		n.mu.Lock()
	}
	n.lastCall = time.Now()
	n.mu.Unlock()
}

// isRateLimitError checks if the error is a 429 rate limit response.
func isRateLimitError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "429")
}

// GetEthClient returns the primary node's eth client (used by governor modules).
func (n *Node) GetEthClient() *ethclient.Client {
	return n.primary.GetEthClient()
}

// GetTxGasFeesAndUsed delegates to the primary node.
func (n *Node) GetTxGasFeesAndUsed(ctx context.Context, tx *evmtypes.Tx) (*big.Int, *big.Int, error) {
	return n.primary.GetTxGasFeesAndUsed(ctx, tx)
}
