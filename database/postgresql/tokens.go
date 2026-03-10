package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dao-portal/extractor/types"
)

// -------------------------------------------------------------------------------------------------------------------
// ---- Tokens table row
// -------------------------------------------------------------------------------------------------------------------

type TokenRow struct {
	ID     types.TokenID `db:"id"`
	Symbol string        `db:"symbol"`
	Name   string        `db:"name"`
}

func (r *TokenRow) ToToken() types.BaseToken {
	baseToken := types.NewBaseToken(r.Symbol, r.Name)
	baseToken.WithID(r.ID)
	return baseToken
}

// -------------------------------------------------------------------------------------------------------------------
// ---- token_natives table row
// -------------------------------------------------------------------------------------------------------------------

type TokenNativeRow struct {
	ToknID       types.TokenID      `db:"token_id"`
	BlockchainID types.BlockchainID `db:"blockchain_id"`
	Denom        string             `db:"denom"`
	Decimals     uint8              `db:"decimals"`
}

func (r *TokenNativeRow) ToNativeToken(baseTokenRow *TokenRow) *types.NativeToken {
	baseToken := baseTokenRow.ToToken()
	return &types.NativeToken{
		BaseToken:    baseToken,
		BlockchainID: r.BlockchainID,
		Denom:        r.Denom,
		Decimals:     r.Decimals,
	}
}

// -------------------------------------------------------------------------------------------------------------------
// ---- tokens_contract table row
// -------------------------------------------------------------------------------------------------------------------

type TokenContractRow struct {
	TokenID           types.TokenID               `db:"token_id"`
	BlockchainID      types.BlockchainID          `db:"blockchain_id"`
	ContractAddressID types.AddressID             `db:"contract_address_id"`
	Standard          types.TokenContractStandard `db:"standard"`
	Decimals          uint8                       `db:"decimals"`
}

func (r *TokenContractRow) ToContractToken(baseTokenRow *TokenRow) *types.ContractToken {
	baseToken := baseTokenRow.ToToken()
	return &types.ContractToken{
		BaseToken:         baseToken,
		BlockchainID:      r.BlockchainID,
		ContractAddressID: r.ContractAddressID,
		Standard:          r.Standard,
		Decimals:          r.Decimals,
	}
}

// -----------------------------------------------------------------------------
// ---- Tokens table functions
// -----------------------------------------------------------------------------

// InsertToken inserts a token into the database.
func (db *DB) InsertToken(ctx context.Context, token types.Token, allowConflict bool) (types.Token, error) {
	// Ensure is a supported token type
	switch token.GetType() {
	case types.TokenTypeNative:
		nativeToken, ok := token.(*types.NativeToken)
		if !ok {
			return nil, fmt.Errorf("invalid token type: %s", token.GetType())
		}
		return db.InsertNativeToken(ctx, nativeToken, allowConflict)
	case types.TokenTypeContract:
		contractToken, ok := token.(*types.ContractToken)
		if !ok {
			return nil, fmt.Errorf("invalid token type: %s", token.GetType())
		}
		return db.InsertContractToken(ctx, contractToken, allowConflict)
	default:
		return nil, fmt.Errorf("unsupported token type: %s", token.GetType())
	}
}

// storeBaseTokenInfo stores the base token information in the database.
func (db *DB) storeBaseTokenInfo(ctx context.Context, tx *sql.Tx, token *types.BaseToken) (TokenRow, error) {
	stmt := `
	WITH ins AS (
		INSERT INTO tokens (symbol, name)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
		RETURNING id, symbol, name
	)
	SELECT * FROM ins
	UNION ALL
	SELECT id, symbol, name
	FROM tokens
	WHERE NOT EXISTS (SELECT 1 FROM ins) AND symbol = $1
	`

	var row TokenRow
	err := tx.QueryRowContext(ctx, stmt, token.GetSymbol(), token.GetName()).
		Scan(&row.ID, &row.Symbol, &row.Name)
	if err != nil {
		return row, err
	}

	return row, nil
}

// InsertNativeToken inserts a native token into the database.
func (db *DB) InsertNativeToken(ctx context.Context, nativeToken *types.NativeToken, allowConflict bool) (*types.NativeToken, error) {
	tx, err := db.SQL.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	// Store the basic token information
	tokendRow, err := db.storeBaseTokenInfo(ctx, tx, &nativeToken.BaseToken)
	if err != nil {
		return nil, fmt.Errorf("failed to store base token information: %w", err)
	}

	// Store the native token information
	stmt := `
	INSERT INTO token_natives(token_id, blockchain_id, denom, decimals) 
	VALUES ($1, $2, $3, $4) 
	RETURNING blockchain_id, denom, decimals
	`
	if allowConflict {
		stmt = `
		WITH ins AS (
			INSERT INTO token_natives (token_id, blockchain_id, denom, decimals)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT DO NOTHING
			RETURNING blockchain_id, denom, decimals
		)
		SELECT * FROM ins
		UNION ALL
		SELECT blockchain_id, denom, decimals
		FROM token_natives
		WHERE NOT EXISTS (SELECT 1 FROM ins) AND token_id = $1 AND blockchain_id = $2 AND denom = $3
		`
	}
	// Insert the token an return the inserted row
	var nativeTokenRow TokenNativeRow
	err = tx.QueryRowContext(ctx, stmt,
		tokendRow.ID,
		nativeToken.BlockchainID,
		nativeToken.Denom,
		nativeToken.Decimals,
	).Scan(&nativeTokenRow.BlockchainID, &nativeTokenRow.Denom, &nativeTokenRow.Decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to insert native token information: %w", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nativeTokenRow.ToNativeToken(&tokendRow), nil
}

// InsertContractToken inserts a contract token into the database.
func (db *DB) InsertContractToken(ctx context.Context, contractToken *types.ContractToken, allowConflict bool) (*types.ContractToken, error) {
	// Start a transaction
	tx, err := db.SQL.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	// Store the basic token information
	tokendRow, err := db.storeBaseTokenInfo(ctx, tx, &contractToken.BaseToken)
	if err != nil {
		return nil, fmt.Errorf("failed to store base token information: %w", err)
	}

	// Store the contract token information
	stmt := `
	INSERT INTO token_contracts(token_id, blockchain_id, contract_address_id, standard, decimals) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING *
	`
	if allowConflict {
		stmt = `
		WITH ins AS (
			INSERT INTO token_contracts(token_id, blockchain_id, contract_address_id, standard, decimals) 
			VALUES ($1, $2, $3, $4, $5) 
			ON CONFLICT DO NOTHING
			RETURNING blockchain_id, contract_address_id, standard, decimals
		)
		SELECT * FROM ins
		UNION ALL
		SELECT blockchain_id, contract_address_id, standard, decimals
		FROM token_contracts
		WHERE NOT EXISTS (SELECT 1 FROM ins) AND token_id = $1 AND blockchain_id = $2 AND contract_address_id = $3
		`
	}

	var contractTokenRow TokenContractRow
	err = tx.QueryRowContext(ctx, stmt,
		tokendRow.ID,
		contractToken.BlockchainID,
		contractToken.ContractAddressID,
		contractToken.Standard,
		contractToken.Decimals,
	).Scan(&contractTokenRow.BlockchainID, &contractTokenRow.ContractAddressID, &contractTokenRow.Standard, &contractTokenRow.Decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to insert contract token information: %w", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return contractTokenRow.ToContractToken(&tokendRow), nil
}

// -------------------------------------------------------------------------------------------------------------------
// ---- Tokens unifed view functions
// -------------------------------------------------------------------------------------------------------------------

type TokenUnifiedViewRow struct {
	TokenID      types.TokenID      `db:"token_id"`
	Symbol       string             `db:"symbol"`
	Name         string             `db:"name"`
	TokenType    types.TokenType    `db:"token_type"`
	BlockChainID types.BlockchainID `db:"blockchain_id"`
	ChainID      types.ChainID      `db:"chain_id"`
	ChainName    string             `db:"chain_name"`
	ChainType    types.ChainType    `db:"chain_type"`
	Decimals     uint8              `db:"decimals"`

	// Present if the token is a native token
	Denom *string `db:"denom"`

	// Present if the token is a contract
	ContractAddressID       *types.AddressID             `db:"contract_address_id"`
	ContractAddress         *string                      `db:"contract_address"`
	ContractAddressEncoding *types.AddressEncoding       `db:"contract_address_encoding"`
	TokenStandard           *types.TokenContractStandard `db:"token_standard"`
}

func (r *TokenUnifiedViewRow) ToToken() types.Token {
	baseToken := types.NewBaseToken(r.Symbol, r.Name)
	baseToken.WithID(r.TokenID)

	switch r.TokenType {
	case types.TokenTypeNative:
		return types.NewNativeToken(baseToken, r.BlockChainID, *r.Denom, r.Decimals)
	case types.TokenTypeContract:
		return types.NewContractToken(baseToken, r.BlockChainID, *r.ContractAddressID, *r.TokenStandard, r.Decimals)
	default:
		panic("unreachable")
	}
}

// getTokensFromUnifiedViewByStmt returns the tokens from the tokens_unified_view view.
func (db *DB) getTokensFromUnifiedViewByStmt(ctx context.Context, stmt string, args ...any) ([]types.Token, error) {
	var rows []TokenUnifiedViewRow
	err := db.SQL.SelectContext(ctx, &rows, stmt, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	tokens := make([]types.Token, len(rows))
	for i, row := range rows {
		tokens[i] = row.ToToken()
	}

	return tokens, nil
}
