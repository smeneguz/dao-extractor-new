package governorbravo

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	fluxevmtypes "github.com/dao-portal/extractor/flux/evm/types"
)

// HandleBlock routes block logs to the correct DAO handler.
func (m *Module) HandleBlock(ctx context.Context, block *fluxevmtypes.Block) error {
	height := uint64(block.GetHeight())

	// Process deferred operations for every DAO.
	for _, inst := range m.daoInstances {
		if err := m.handleDeferredOperations(ctx, block, inst); err != nil {
			return fmt.Errorf("deferred ops for %s: %w", inst.DAO.Name, err)
		}
	}

	// Route each log entry to its matching DAO.
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

			if err := m.handleGovernorBravoEvents(ctx, block, tx, logEntry, inst); err != nil {
				return fmt.Errorf("governor-bravo events for %s: %w", inst.DAO.Name, err)
			}
		}
	}

	return nil
}
