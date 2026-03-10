package postgresql

import (
	"context"

	"github.com/dao-portal/extractor/types"
)

// StoreTokenTransfer persists an ERC20 Transfer event to the database.
func (db *DB) StoreTokenTransfer(ctx context.Context, t *types.TokenTransfer) error {
	stmt := `
	INSERT INTO token_transfers(
		dao_id, chain_id, token_address,
		from_address, to_address, amount,
		tx_hash, block_height, log_index, ts)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	ON CONFLICT (tx_hash, log_index) DO NOTHING
	`
	_, err := db.SQL.ExecContext(ctx, stmt,
		t.DaoID,
		t.ChainID,
		t.TokenAddress,
		t.FromAddress,
		t.ToAddress,
		t.Amount.String(),
		t.TxHash,
		t.BlockHeight,
		t.LogIndex,
		t.Timestamp.UTC(),
	)
	return err
}

// StoreDelegationEvent persists a DelegateChanged event to the database.
func (db *DB) StoreDelegationEvent(ctx context.Context, d *types.DelegationEvent) error {
	stmt := `
	INSERT INTO delegation_events(
		dao_id, chain_id, token_address,
		delegator, from_delegate, to_delegate,
		tx_hash, block_height, log_index, ts)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	ON CONFLICT (tx_hash, log_index) DO NOTHING
	`
	_, err := db.SQL.ExecContext(ctx, stmt,
		d.DaoID,
		d.ChainID,
		d.TokenAddress,
		d.Delegator,
		d.FromDelegate,
		d.ToDelegate,
		d.TxHash,
		d.BlockHeight,
		d.LogIndex,
		d.Timestamp.UTC(),
	)
	return err
}

// StoreDelegateVotesChanged persists a DelegateVotesChanged event to the database.
func (db *DB) StoreDelegateVotesChanged(ctx context.Context, d *types.DelegateVotesChanged) error {
	stmt := `
	INSERT INTO delegate_votes_changed(
		dao_id, chain_id, token_address,
		delegate, previous_balance, new_balance,
		tx_hash, block_height, log_index, ts)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	ON CONFLICT (tx_hash, log_index) DO NOTHING
	`
	_, err := db.SQL.ExecContext(ctx, stmt,
		d.DaoID,
		d.ChainID,
		d.TokenAddress,
		d.Delegate,
		d.PreviousBalance.String(),
		d.NewBalance.String(),
		d.TxHash,
		d.BlockHeight,
		d.LogIndex,
		d.Timestamp.UTC(),
	)
	return err
}
