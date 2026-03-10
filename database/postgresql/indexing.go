package postgresql

import (
	"context"
	"encoding/json"

	dbtypes "github.com/dao-portal/extractor/database/postgresql/types"
	"github.com/dao-portal/extractor/types"
)

type HeightDeferredOperationRow struct {
	CreatorKey string           `db:"creator_key"`
	Type       string           `db:"type"`
	Height     dbtypes.BigInt   `db:"height"`
	Payload    *json.RawMessage `db:"payload"`
}

func (r *HeightDeferredOperationRow) ToOperation() *types.HeightDeferredOperation {
	var payload json.RawMessage
	if r.Payload != nil && len(*r.Payload) > 0 {
		payload = *r.Payload
	}
	return types.NewHeightDeferredOperation(
		r.CreatorKey,
		r.Type,
		r.Height.Int,
		payload,
	)
}

// StoreHeightDeferredOperation stores a new operation that should be executed at a certain height.
func (db *DB) StoreHeightDeferredOperation(ctx context.Context, op *types.HeightDeferredOperation) error {
	var payload *json.RawMessage
	if len(op.Payload) > 0 {
		payload = &op.Payload
	}

	_, err := db.SQL.ExecContext(ctx, `
		INSERT INTO height_deferred_operations (creator_key, height, type, payload)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (creator_key, height, type) DO UPDATE SET payload = $4`,
		op.CreatorKey, op.Height.String(), op.Type, payload)

	return err
}

func (db *DB) getHeightDeferredOperationsByStmt(
	ctx context.Context, stmt string, args ...any,
) ([]types.HeightDeferredOperation, error) {
	var rows []HeightDeferredOperationRow
	err := db.SQL.SelectContext(ctx, &rows, stmt, args...)
	if err != nil {
		return nil, err
	}

	ops := make([]types.HeightDeferredOperation, len(rows))
	for i, row := range rows {
		decoedOp := row.ToOperation()
		ops[i] = *decoedOp
	}

	return ops, nil
}

func (db *DB) GetHeightDeferredOperations(
	ctx context.Context, creatorKey string, opType string,
) ([]types.HeightDeferredOperation, error) {
	return db.getHeightDeferredOperationsByStmt(ctx, `
		SELECT *
		FROM height_deferred_operations
		WHERE creator_key = $1 AND type = $2
		ORDER BY height ASC
		`,
		creatorKey, opType)
}

func (db *DB) RemoveHeightDeferredOperations(ctx context.Context, op *types.HeightDeferredOperation) error {
	_, err := db.SQL.ExecContext(ctx, `
		DELETE FROM height_deferred_operations
		WHERE creator_key = $1 AND type = $2 AND height = $3
		`,
		op.CreatorKey, op.Type, op.Height.Text(10))

	return err
}
