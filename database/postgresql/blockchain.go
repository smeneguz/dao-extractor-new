// This file contains the code to interact with the tables present inside the 'schema/01-blockchain.sql' file.
package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dao-portal/extractor/types"
)

// -------------------------------------------------------------------------------------------------------------------
// ---- Blockchains table functions
// -------------------------------------------------------------------------------------------------------------------

// BlockChainRow represents a row from the blockchains table.
type BlockChainRow struct {
	ID      types.BlockchainID `db:"id"`
	ChainID types.ChainID      `db:"chain_id"`
	Name    string             `db:"name"`
	Type    types.ChainType    `db:"type"`
}

func (r *BlockChainRow) ToBlockchain() *types.Blockchain {
	return types.NewBlockchain(r.ChainID, r.Name, r.Type).WithID(r.ID)
}

// InsertBlockchain inserts a new blockchain into the database.
func (db *DB) InsertBlockchain(
	ctx context.Context,
	blockchain *types.Blockchain,
	allowConflict bool,
) (*types.Blockchain, error) {
	stmt := `INSERT INTO blockchains(chain_id, name, type) VALUES ($1, $2, $3) RETURNING id`
	if allowConflict {
		stmt = `
		WITH inserted AS (
			INSERT INTO blockchains (chain_id, name, type)
			VALUES ($1, $2, $3)
			ON CONFLICT (chain_id) DO NOTHING
			RETURNING id
		)
		SELECT id FROM inserted
		UNION ALL
		SELECT id FROM blockchains WHERE chain_id = $1
		LIMIT 1;
		`
	}

	var id types.BlockchainID
	err := db.SQL.QueryRowContext(ctx, stmt,
		blockchain.ChainID,
		blockchain.Name,
		blockchain.Type,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return types.NewBlockchain(blockchain.ChainID, blockchain.Name, blockchain.Type).WithID(id), nil
}

// UpdateBlockchain updates an existing blockchain in the database.
// The Blockchain object must exist and have the same ID and ChainID as the one provided.
func (db *DB) UpdateBlockchain(ctx context.Context, blockchain *types.Blockchain) (*types.Blockchain, error) {
	var id types.BlockchainID
	err := db.SQL.QueryRowContext(ctx, `
		UPDATE blockchains
		SET name = $2, type = $3
		WHERE chain_id = $1
		RETURNING id
	`, blockchain.ChainID,
		blockchain.Name,
		blockchain.Type,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return types.NewBlockchain(blockchain.ChainID, blockchain.Name, blockchain.Type).WithID(id), nil
}

// GetBlockchainByStmt returns a blockchain by executing the provided SQL statement.
func (db *DB) getBlockchainByStmt(ctx context.Context, stmt string, args ...any) (*types.Blockchain, error) {
	var row BlockChainRow
	err := db.SQL.GetContext(ctx, &row, stmt, args...)
	// Handle the case where the blockchain is not found
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return row.ToBlockchain(), nil
}

// GetBlockchain returns the blockchain with the given chain ID.
// If the blockchain is not found, it returns nil.
func (db *DB) GetBlockchainByChainID(ctx context.Context, chainID string) (*types.Blockchain, error) {
	return db.getBlockchainByStmt(ctx, `
		SELECT *
		FROM blockchains
		WHERE chain_id = $1
	`, chainID)
}

// -------------------------------------------------------------------------------------------------------------------
// ---- Addresses table functions
// -------------------------------------------------------------------------------------------------------------------

// AddressRow represents a row from the addresses table.
type AddressRow struct {
	ID         types.AddressID       `db:"id"`
	Address    string                `db:"address"`
	Label      string                `db:"label"`
	IsContract bool                  `db:"is_contract"`
	Encoding   types.AddressEncoding `db:"encoding"`
}

func (r *AddressRow) ToAddress() *types.Address {
	return types.NewAddress(r.Address, r.Label, r.IsContract, r.Encoding).WithID(r.ID)
}

// InsertAddress inserts a new address into the database.
func (db *DB) InsertAddress(ctx context.Context, address *types.Address, allowConflict bool) (*types.Address, error) {
	stmt := `INSERT INTO addresses(address, label, is_contract, encoding) VALUES ($1, $2, $3, $4) RETURNING id`
	if allowConflict {
		stmt = `
		WITH inserted AS (
			INSERT INTO addresses(address, label, is_contract, encoding)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (address) DO NOTHING
			RETURNING id
		)
		SELECT id FROM inserted
		UNION ALL
		SELECT id FROM addresses WHERE address = $1
		LIMIT 1;
		`
	}

	var id types.AddressID
	err := db.SQL.QueryRowContext(ctx, stmt,
		address.Address,
		address.Label,
		address.IsContract,
		address.Encoding,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return types.NewAddress(address.Address, address.Label, address.IsContract, address.Encoding).WithID(id), nil
}

// UpdateAddress updates an existing address in the database.
// The Address object must exist and have the same ID and Address as the one provided.
func (db *DB) UpdateAddress(ctx context.Context, address *types.Address) (*types.Address, error) {
	var id types.AddressID
	err := db.SQL.QueryRowContext(ctx, `
		UPDATE addresses
		SET label = $2, is_contract = $3, encoding = $4
		WHERE address = $1
		RETURNING id
	`, address.Address,
		address.Label,
		address.IsContract,
		address.Encoding,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return types.NewAddress(address.Address, address.Label, address.IsContract, address.Encoding).WithID(id), nil
}

func (db *DB) getAddressByStmt(ctx context.Context, stmt string, args ...any) (*types.Address, error) {
	var row AddressRow
	err := db.SQL.GetContext(ctx, &row, stmt, args...)
	// Handle the case where the address is not found
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return row.ToAddress(), nil
}

// GetAddressByAddress returns the address with the given address.
// If the address is not found, it returns nil.
func (db *DB) GetAddressByAddress(ctx context.Context, address string) (*types.Address, error) {
	return db.getAddressByStmt(ctx, `
		SELECT *
		FROM addresses
		WHERE address = $1
	`, address)
}

func (db *DB) getAddressesByStmt(ctx context.Context, stmt string, args ...any) ([]*types.Address, error) {
	var rows []AddressRow
	err := db.SQL.SelectContext(ctx, &rows, stmt, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	addresses := make([]*types.Address, len(rows))
	for i, row := range rows {
		addresses[i] = row.ToAddress()
	}
	return addresses, nil
}

// -------------------------------------------------------------------------------------------------------------------
// ---- Address chains association table functions
// -------------------------------------------------------------------------------------------------------------------

// AddressChainRow represents a row from the address_blockchains table.
type AddressChainRow struct {
	AddressID    types.AddressID    `db:"address_id"`
	BlockChainID types.BlockchainID `db:"blockchain_id"`
}

// AssociateAddressToBlockchain associates an address to a blockchain.
func (db *DB) AssociateAddressToBlockchain(ctx context.Context, address *types.Address, blockchain *types.Blockchain) error {
	_, err := db.SQL.ExecContext(ctx, `
		INSERT INTO address_blockchains(address_id, blockchain_id)
		VALUES ($1, $2)
		ON CONFLICT (address_id, blockchain_id) DO NOTHING
	`, address.ID, blockchain.ID)
	return err
}
