package postgresql

import (
	"context"
	"encoding/json"

	"github.com/dao-portal/extractor/types"
)

// StoreRawEvent persists a raw EVM event log to the database.
func (db *DB) StoreRawEvent(ctx context.Context, event *types.RawEvent) error {
	topicsJSON, err := json.Marshal(event.Topics)
	if err != nil {
		return err
	}

	stmt := `
	INSERT INTO raw_events(
		dao_id, chain_id, contract_address_id,
		tx_hash, block_height, log_index, ts,
		topics, data)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	ON CONFLICT (tx_hash, log_index) DO NOTHING
	`
	_, err = db.SQL.ExecContext(ctx, stmt,
		event.DaoID,
		event.ChainID,
		event.ContractAddressID,
		event.TxHash,
		event.BlockHeight,
		event.LogIndex,
		event.Timestamp.UTC(),
		topicsJSON,
		event.Data,
	)
	return err
}
