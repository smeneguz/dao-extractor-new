package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dao-portal/extractor/types"
)

// -------------------------------------------------------------------------------------------------------------------
// ---- DAOs table functions
// -------------------------------------------------------------------------------------------------------------------

// DAORow represents a row from the daos table.
type DAORow struct {
	ID     types.DAOID     `db:"id"`
	Symbol types.DAOSymbol `db:"symbol"`
	Name   string          `db:"name"`
}

func (r *DAORow) ToDAO() *types.DAO {
	return types.NewDAO(r.Symbol, r.Name).WithID(r.ID)
}

// InsertDAO inserts a new DAO into the database.
func (db *DB) InsertDAO(ctx context.Context, dao *types.DAO, allowConflict bool) (*types.DAO, error) {
	stmt := `INSERT INTO daos(symbol, name) VALUES ($1, $2) RETURNING id`
	if allowConflict {
		stmt = `
		WITH inserted AS (
			INSERT INTO daos(symbol, name)
			VALUES ($1, $2)
			ON CONFLICT (symbol) DO NOTHING
			RETURNING id
		)
		SELECT id FROM inserted
		UNION ALL
		SELECT id FROM daos WHERE symbol = $1
		LIMIT 1;
		`
	}

	var id types.DAOID
	err := db.SQL.QueryRowContext(ctx, stmt,
		dao.Symbol,
		dao.Name,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return types.NewDAO(dao.Symbol, dao.Name).WithID(id), nil
}

// UpdateDAO updates an existing DAO in the database.
// The DAO object must exist and have the same ID and Symbol as the one provided.
func (db *DB) UpdateDAO(ctx context.Context, dao *types.DAO) (*types.DAO, error) {
	var id types.DAOID
	err := db.SQL.QueryRowContext(ctx, `
		UPDATE daos
		SET name = $2
		WHERE symbol = $1
		RETURNING id
	`, dao.Symbol,
		dao.Name,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return types.NewDAO(dao.Symbol, dao.Name).WithID(id), nil
}

func (db *DB) getDAOByStmt(ctx context.Context, stmt string, args ...any) (*types.DAO, error) {
	var row DAORow
	err := db.SQL.GetContext(ctx, &row, stmt, args...)
	// Handle the case where the DAO is not found
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return row.ToDAO(), nil
}

// GetDAOBySymbol returns the DAO with the given symbol.
// If the DAO is not found, it returns nil.
func (db *DB) GetDAOBySymbol(ctx context.Context, symbol string) (*types.DAO, error) {
	return db.getDAOByStmt(ctx, `
		SELECT *
		FROM daos
		WHERE symbol = $1
	`, symbol)
}

// -------------------------------------------------------------------------------------------------------------------
// ---- DAO contracts association functions
// -------------------------------------------------------------------------------------------------------------------
type DaoContractRow struct {
	DaoID             uint64 `db:"dao_id"`
	ContractAddressID uint64 `db:"contract_address_id"`
	BlockchainID      uint64 `db:"blockchain_id"`
}

// AssociateDAOToContract associates a DAO to a contract and the blockchain where the contract is deployed.
func (db *DB) AssociateDAOToContract(
	ctx context.Context, dao *types.DAO, contractAddress *types.Address, blockchain *types.Blockchain,
) error {
	if dao == nil {
		return fmt.Errorf("dao must not be nil")
	}
	if contractAddress == nil || !contractAddress.IsContract {
		return fmt.Errorf("contract address must be a contract")
	}
	if blockchain == nil {
		return fmt.Errorf("blockchain must not be nil")
	}

	_, err := db.SQL.ExecContext(ctx, `
		INSERT INTO dao_contracts(dao_id, contract_address_id, blockchain_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (dao_id, contract_address_id, blockchain_id) DO NOTHING
	`, dao.ID, contractAddress.ID, blockchain.ID)
	return err
}

// GetDaoContractsOnChain returns the contracts that are associated to a DAO on a specific blockchain.
func (db *DB) GetDaoContractsOnChain(
	ctx context.Context, dao *types.DAO, blockchain *types.Blockchain, pagination *types.Pagination,
) ([]*types.Address, error) {
	if pagination == nil {
		pagination = db.defaultPagination
	}

	return db.getAddressesByStmt(ctx, `
		SELECT a.*
		FROM dao_contracts dc
		JOIN addresses a ON a.id = dc.contract_address_id
		WHERE dao_id = $1 AND blockchain_id = $2
		ORDER BY a.id
		LIMIT $3 OFFSET $4
	`, dao.ID, blockchain.ID, pagination.Limit, pagination.Offset)
}

// -------------------------------------------------------------------------------------------------------------------
// ---- DAO treasury tokens association functions
// -------------------------------------------------------------------------------------------------------------------

type DaoTokenRow struct {
	DaoID   types.DAOID   `db:"dao_id"`
	TokenID types.TokenID `db:"token_id"`
}

// AssociateDAOToToken associates a DAO to a token.
func (db *DB) SetDAOTreasuryToken(ctx context.Context, dao *types.DAO, token types.Token) error {
	_, err := db.SQL.ExecContext(ctx, `
		INSERT INTO dao_treasury_tokens(dao_id, token_id)
		VALUES ($1, $2)
		ON CONFLICT (dao_id, token_id) DO NOTHING
	`, dao.ID, token.GetID())
	return err
}
