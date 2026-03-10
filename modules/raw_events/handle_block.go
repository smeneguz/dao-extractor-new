package rawevents

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	fluxevmtypes "github.com/dao-portal/extractor/flux/evm/types"
	"github.com/dao-portal/extractor/types"
)

// HandleBlock routes block logs to the correct DAO and stores raw events.
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

			// Resolve the persisted address for this contract.
			persistedAddr, ok := inst.ContractAddresses[addr]
			if !ok || persistedAddr == nil {
				continue
			}

			// Build topics list from the log entry.
			topics := make([]string, 0, len(logEntry.Topics))
			for _, t := range logEntry.Topics {
				topics = append(topics, t.Hex())
			}

			rawEvent := &types.RawEvent{
				DaoID:             inst.DAO.ID,
				ChainID:          inst.Chain.ID,
				ContractAddressID: persistedAddr.ID,
				TxHash:           tx.Hash.NormalizedHex(),
				BlockHeight:      height,
				LogIndex:         logEntry.Index,
				Timestamp:        timestamp,
				Topics:           topics,
				Data:             common.Bytes2Hex(logEntry.Data),
			}

			if err := m.db.StoreRawEvent(ctx, rawEvent); err != nil {
				return fmt.Errorf("store raw event for %s: %w", inst.DAO.Name, err)
			}
		}
	}

	return nil
}
