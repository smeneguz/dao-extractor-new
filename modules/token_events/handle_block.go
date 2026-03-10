package tokenevents

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	fluxevmtypes "github.com/dao-portal/extractor/flux/evm/types"
	"github.com/dao-portal/extractor/types"
)

// HandleBlock routes block logs to the correct DAO and stores token events.
func (m *Module) HandleBlock(ctx context.Context, block *fluxevmtypes.Block) error {
	height := uint64(block.GetHeight())
	timestamp := block.GetTimeStamp()

	for _, tx := range block.Txs {
		for i := range tx.Logs {
			logEntry := &tx.Logs[i]
			addr := common.BytesToAddress(logEntry.Address.TrimLeadingZero())

			inst, ok := m.addressToDAO[addr]
			if !ok || inst == nil {
				continue
			}

			// Skip logs from before the contract's deployment block.
			if inst.DeployedAt > 0 && height < inst.DeployedAt {
				continue
			}

			// Must have at least topic0 (event signature).
			if len(logEntry.Topics) == 0 {
				continue
			}

			// Convert topic0 from EVMBytes to common.Hash for event identification.
			topic0 := common.BytesToHash(logEntry.Topics[0])
			eventName := identifyEvent(topic0)
			if eventName == "" {
				continue
			}

			// Convert all flux topics to go-ethereum common.Hash.
			topics := make([]common.Hash, len(logEntry.Topics))
			for j, t := range logEntry.Topics {
				topics[j] = common.BytesToHash(t)
			}

			tokenAddr, hasToken := m.addressToTokenStr[addr]
			if !hasToken || tokenAddr == "" {
				m.logger.Warn().
					Str("addr", addr.Hex()).
					Uint64("height", height).
					Msg("token address not found in routing map, skipping event")
				continue
			}
			txHash := tx.Hash.NormalizedHex()
			chainID := string(inst.Chain.ChainID)

			switch eventName {
			case eventNameTransfer:
				from, to, value, ok := parseTransferEvent(topics, []byte(logEntry.Data))
				if !ok {
					m.logger.Warn().
						Str("dao", inst.DAO.Name).
						Str("tx", txHash).
						Uint64("height", height).
						Msg("failed to parse Transfer event")
					continue
				}
				err := m.db.StoreTokenTransfer(ctx, &types.TokenTransfer{
					DaoID:        inst.DAO.ID,
					ChainID:      chainID,
					TokenAddress: tokenAddr,
					FromAddress:  from.Hex(),
					ToAddress:    to.Hex(),
					Amount:       value,
					TxHash:       txHash,
					BlockHeight:  height,
					LogIndex:     logEntry.Index,
					Timestamp:    timestamp,
				})
				if err != nil {
					return fmt.Errorf("store Transfer for %s: %w", inst.DAO.Name, err)
				}

			case eventNameDelegateChanged:
				delegator, fromDelegate, toDelegate, ok := parseDelegateChangedEvent(topics)
				if !ok {
					m.logger.Warn().
						Str("dao", inst.DAO.Name).
						Str("tx", txHash).
						Uint64("height", height).
						Msg("failed to parse DelegateChanged event")
					continue
				}
				err := m.db.StoreDelegationEvent(ctx, &types.DelegationEvent{
					DaoID:        inst.DAO.ID,
					ChainID:      chainID,
					TokenAddress: tokenAddr,
					Delegator:    delegator.Hex(),
					FromDelegate: fromDelegate.Hex(),
					ToDelegate:   toDelegate.Hex(),
					TxHash:       txHash,
					BlockHeight:  height,
					LogIndex:     logEntry.Index,
					Timestamp:    timestamp,
				})
				if err != nil {
					return fmt.Errorf("store DelegateChanged for %s: %w", inst.DAO.Name, err)
				}

			case eventNameDelegateVotesChanged:
				delegate, prevBalance, newBalance, ok := parseDelegateVotesChangedEvent(topics, []byte(logEntry.Data))
				if !ok {
					m.logger.Warn().
						Str("dao", inst.DAO.Name).
						Str("tx", txHash).
						Uint64("height", height).
						Msg("failed to parse DelegateVotesChanged event")
					continue
				}
				err := m.db.StoreDelegateVotesChanged(ctx, &types.DelegateVotesChanged{
					DaoID:           inst.DAO.ID,
					ChainID:         chainID,
					TokenAddress:    tokenAddr,
					Delegate:        delegate.Hex(),
					PreviousBalance: prevBalance,
					NewBalance:      newBalance,
					TxHash:          txHash,
					BlockHeight:     height,
					LogIndex:        logEntry.Index,
					Timestamp:       timestamp,
				})
				if err != nil {
					return fmt.Errorf("store DelegateVotesChanged for %s: %w", inst.DAO.Name, err)
				}
			}
		}
	}

	return nil
}
