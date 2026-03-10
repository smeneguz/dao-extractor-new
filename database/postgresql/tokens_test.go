package postgresql_test

import (
	"context"

	"github.com/dao-portal/extractor/database/postgresql"
	"github.com/dao-portal/extractor/types"
)

func (suite *DbTestSuite) TestInsertNativeToken() {
	testCases := []struct {
		name          string
		setup         func(ctx context.Context)
		token         types.Token
		allowConflict bool
		shouldErr     bool
		check         func()
	}{
		{
			name: "insert a native token referencing a non existing blockchain fails",
			token: types.NewNativeToken(
				types.NewBaseToken("ATOM", "Cosmos Hub Token"),
				1,
				"uatom",
				6,
			),
			shouldErr: true,
		},
		{
			name: "insert a native token referencing an existing blockchain correctly",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
			},
			token: types.NewNativeToken(
				types.NewBaseToken("ETH", "Ethereum"),
				1,
				"eth",
				18,
			),
			check: func() {
				var baseToken postgresql.TokenRow
				err := suite.database.SQL.Get(&baseToken,
					`SELECT * FROM tokens WHERE symbol = $1`,
					"ETH",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.TokenID(1), baseToken.ID)
				suite.Require().Equal("ETH", baseToken.Symbol)
				suite.Require().Equal("Ethereum", baseToken.Name)

				var token postgresql.TokenNativeRow
				err = suite.database.SQL.Get(&token,
					`SELECT * FROM token_natives WHERE token_id = $1 AND blockchain_id = $2 AND denom = $3`,
					baseToken.ID, 1, "eth",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.BlockchainID(1), token.BlockchainID)
				suite.Require().Equal("eth", token.Denom)
				suite.Require().Equal(uint8(18), token.Decimals)
			},
		},
		{
			name: "insert a conflicting native token fails",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)

				// Insert the token
				_, err = suite.database.SQL.ExecContext(ctx, `
				INSERT INTO tokens(id, symbol, name) 
				VALUES ($1, $2, $3) 
				`,
					1, "ETH", "Ethereum",
				)
				suite.Require().NoError(err)

				// Insert the native token information
				_, err = suite.database.SQL.ExecContext(ctx, `
				INSERT INTO token_natives(token_id, blockchain_id, denom, decimals) 
				VALUES ($1, $2, $3, $4) 
				`,
					1, 1, "eth", 18,
				)
				suite.Require().NoError(err)
			},
			token: types.NewNativeToken(
				types.NewBaseToken("ETH", "Ethereum 2"),
				1,
				"eth",
				18,
			),
			allowConflict: false,
			shouldErr:     true,
		},
		{
			name: "insert a native token with allow conflict don not update the token",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)

				// Insert the token
				_, err = suite.database.SQL.ExecContext(ctx, `
				INSERT INTO tokens(id, symbol, name) 
				VALUES ($1, $2, $3) 
				`,
					1, "ETH", "Ethereum",
				)
				suite.Require().NoError(err)

				// Insert the native token information
				_, err = suite.database.SQL.ExecContext(ctx, `
				INSERT INTO token_natives(token_id, blockchain_id, denom, decimals) 
				VALUES ($1, $2, $3, $4) 
				`,
					1, 1, "eth", 18,
				)
				suite.Require().NoError(err)
			},
			token: types.NewNativeToken(
				types.NewBaseToken("ETH", "Ethereum 2"),
				1,
				"eth",
				19,
			),
			allowConflict: true,
			check: func() {
				var baseToken postgresql.TokenRow
				err := suite.database.SQL.Get(&baseToken,
					`SELECT * FROM tokens WHERE symbol = $1`,
					"ETH",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.TokenID(1), baseToken.ID)
				suite.Require().Equal("ETH", baseToken.Symbol)
				suite.Require().Equal("Ethereum", baseToken.Name)

				var token postgresql.TokenNativeRow
				err = suite.database.SQL.Get(&token,
					`SELECT * FROM token_natives WHERE token_id = $1 AND blockchain_id = $2 AND denom = $3`,
					baseToken.ID, 1, "eth",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.BlockchainID(1), token.BlockchainID)
				suite.Require().Equal("eth", token.Denom)
				suite.Require().Equal(uint8(18), token.Decimals)
			},
		},
		{
			name: "insert a native token correctly",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
			},
			token: types.NewNativeToken(
				types.NewBaseToken("ETH", "Ethereum"),
				1,
				"eth",
				18,
			),
			allowConflict: false,
			shouldErr:     false,
			check: func() {
				var baseToken postgresql.TokenRow
				err := suite.database.SQL.Get(&baseToken,
					`SELECT * FROM tokens WHERE symbol = $1`,
					"ETH",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.TokenID(1), baseToken.ID)
				suite.Require().Equal("ETH", baseToken.Symbol)
				suite.Require().Equal("Ethereum", baseToken.Name)

				var token postgresql.TokenNativeRow
				err = suite.database.SQL.Get(&token,
					`SELECT * FROM token_natives WHERE token_id = $1`,
					baseToken.ID,
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.BlockchainID(1), token.BlockchainID)
				suite.Require().Equal("eth", token.Denom)
				suite.Require().Equal(uint8(18), token.Decimals)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			_, err := suite.database.InsertToken(suite.T().Context(), tc.token, tc.allowConflict)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				if tc.check != nil {
					tc.check()
				}
			}
		})
	}
}

func (suite *DbTestSuite) TestInsertContractToken() {
	testCases := []struct {
		name          string
		setup         func(ctx context.Context)
		token         types.Token
		allowConflict bool
		shouldErr     bool
		check         func()
	}{
		{
			name: "insert a contract token referencing a non existing blockchain fails",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertAddress(ctx, types.NewAddress("0x1", "label", true, types.AddressEncodingTypeHex), false)
				suite.Require().NoError(err)
			},
			token: types.NewContractToken(
				types.NewBaseToken("ATOM", "Cosmos Hub Token"),
				1,
				1,
				types.TokenContractStandardERC20,
				6,
			),
			shouldErr: true,
		},
		{
			name: "insert a contract token referencing a non existing address fails",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
			},
			token: types.NewContractToken(
				types.NewBaseToken("ATOM", "Cosmos Hub Token"),
				1,
				1,
				types.TokenContractStandardERC20,
				6,
			),
			shouldErr: true,
		},
		{
			name: "insert a conflicting contract token fails",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x1", "USDT ERC20", true, types.AddressEncodingTypeHex), false)
				suite.Require().NoError(err)

				// Insert the token
				_, err = suite.database.SQL.ExecContext(ctx, `
				INSERT INTO tokens(id, symbol, name) 
				VALUES ($1, $2, $3) 
				`,
					1, "USDT", "USDT",
				)
				suite.Require().NoError(err)

				// Insert the contract token information
				_, err = suite.database.SQL.ExecContext(ctx, `
				INSERT INTO token_contracts(token_id, blockchain_id, contract_address_id, standard, decimals)
				VALUES ($1, $2, $3, $4, $5) 
				`,
					1, 1, 1, types.TokenContractStandardERC20, 18,
				)
				suite.Require().NoError(err)
			},
			token: types.NewContractToken(
				types.NewBaseToken("ETH", "Ethereum 2"),
				1,
				1,
				types.TokenContractStandardERC20,
				18,
			),
			allowConflict: false,
			shouldErr:     true,
		},
		{
			name: "insert a contract token with allow conflict don not update the token",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x1", "USDT ERC20", true, types.AddressEncodingTypeHex), false)
				suite.Require().NoError(err)

				// Insert the token
				_, err = suite.database.SQL.ExecContext(ctx, `
				INSERT INTO tokens(id, symbol, name) 
				VALUES ($1, $2, $3) 
				`,
					1, "USDT", "Tether USD",
				)
				suite.Require().NoError(err)

				// Insert the ContractToken information
				_, err = suite.database.SQL.ExecContext(ctx, `
				INSERT INTO token_contracts(token_id, blockchain_id, contract_address_id, standard, decimals)
				VALUES ($1, $2, $3, $4, $5) 
				`,
					1, 1, 1, types.TokenContractStandardERC20, 18,
				)
				suite.Require().NoError(err)
			},
			token: types.NewContractToken(
				types.NewBaseToken("USDT", "USD Tether"),
				1,
				1,
				types.TokenContractStandardERC20,
				19,
			),
			allowConflict: true,
			check: func() {
				var baseToken postgresql.TokenRow
				err := suite.database.SQL.Get(&baseToken,
					`SELECT * FROM tokens WHERE symbol = $1`,
					"USDT",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.TokenID(1), baseToken.ID)
				suite.Require().Equal("USDT", baseToken.Symbol)
				suite.Require().Equal("Tether USD", baseToken.Name)

				var token postgresql.TokenContractRow
				err = suite.database.SQL.Get(&token,
					`SELECT * FROM token_contracts WHERE token_id = $1 AND blockchain_id = $2 AND contract_address_id = $3`,
					baseToken.ID, 1, 1,
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.BlockchainID(1), token.BlockchainID)
				suite.Require().Equal(types.AddressID(1), token.ContractAddressID)
				suite.Require().Equal(types.TokenContractStandardERC20, token.Standard)
				suite.Require().Equal(uint8(18), token.Decimals)
			},
		},
		{
			name: "insert a contract token successfully",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x1", "USDT ERC20", true, types.AddressEncodingTypeHex), false)
				suite.Require().NoError(err)
			},
			token: types.NewContractToken(
				types.NewBaseToken("USDT", "Tether USD"),
				1,
				1,
				types.TokenContractStandardERC20,
				18,
			),
			allowConflict: true,
			check: func() {
				var baseToken postgresql.TokenRow
				err := suite.database.SQL.Get(&baseToken,
					`SELECT * FROM tokens WHERE symbol = $1`,
					"USDT",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.TokenID(1), baseToken.ID)
				suite.Require().Equal("USDT", baseToken.Symbol)
				suite.Require().Equal("Tether USD", baseToken.Name)

				var token postgresql.TokenContractRow
				err = suite.database.SQL.Get(&token,
					`SELECT * FROM token_contracts WHERE token_id = $1 AND blockchain_id = $2 AND contract_address_id = $3`,
					baseToken.ID, 1, 1,
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.BlockchainID(1), token.BlockchainID)
				suite.Require().Equal(types.AddressID(1), token.ContractAddressID)
				suite.Require().Equal(types.TokenContractStandardERC20, token.Standard)
				suite.Require().Equal(uint8(18), token.Decimals)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			_, err := suite.database.InsertToken(suite.T().Context(), tc.token, tc.allowConflict)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				if tc.check != nil {
					tc.check()
				}
			}
		})
	}
}
