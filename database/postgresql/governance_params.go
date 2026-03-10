package postgresql

import (
	"context"

	"github.com/dao-portal/extractor/types"
)

// StoreGovernanceParamChange persists a governance parameter change event.
func (db *DB) StoreGovernanceParamChange(ctx context.Context, p *types.GovernanceParamChange) error {
	stmt := `
	INSERT INTO governance_param_changes(
		dao_id, chain_id, contract_address,
		param_name, old_value, new_value,
		tx_hash, block_height, log_index, ts)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	ON CONFLICT (tx_hash, log_index) DO NOTHING
	`
	_, err := db.SQL.ExecContext(ctx, stmt,
		p.DaoID,
		p.ChainID,
		p.ContractAddress,
		p.ParamName,
		p.OldValue.String(),
		p.NewValue.String(),
		p.TxHash,
		p.BlockHeight,
		p.LogIndex,
		p.Timestamp.UTC(),
	)
	return err
}
