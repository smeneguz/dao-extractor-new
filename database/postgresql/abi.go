package postgresql

import (
	"context"
	"database/sql"
	"errors"

	evmtypes "github.com/dao-portal/extractor/types/evm"
)

func (db *DB) SaveAbi(ctx context.Context, abi *evmtypes.Abi) error {
	_, err := db.SQL.ExecContext(ctx, `
		INSERT INTO abis
		(
			contract_address,
			chain_id,
			abi
		)
		VALUES
		(
			$1, $2, $3
		)
		ON CONFLICT (contract_address, chain_id)
		DO UPDATE SET
			abi = excluded.abi
	`, abi.ContractAddress,
		abi.ChainID,
		abi.ABI,
	)

	return err
}

// GetAbi returns the ABI of the contract with the given address and chain ID.
// If the ABI is not found, it returns nil.
func (db *DB) GetAbi(ctx context.Context, chainID string, contractAddress string) (*evmtypes.Abi, error) {
	var abi evmtypes.Abi

	row := db.SQL.QueryRowContext(ctx, `
		SELECT
			contract_address,
			chain_id,
			abi
		FROM abis
		WHERE contract_address = $1 AND chain_id = $2
	`, contractAddress,
		chainID,
	)
	err := row.Scan(&abi.ContractAddress, &abi.ChainID, &abi.ABI)
	// Handle the case where the ABI is not found
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &abi, nil
}
