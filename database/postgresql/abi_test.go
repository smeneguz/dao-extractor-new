package postgresql_test

import (
	"context"

	evmtypes "github.com/dao-portal/extractor/types/evm"
)

func (suite *DbTestSuite) TestSaveABI() {
	testCases := []struct {
		name      string
		setup     func(ctx context.Context)
		abi       *evmtypes.Abi
		shouldErr bool
		check     func()
	}{
		{
			name:      "save correctly",
			abi:       evmtypes.NewAbi("test", "save_contract", []byte(`{"name": "test"}`)),
			shouldErr: false,
			check: func() {
				var abi evmtypes.Abi
				row := suite.database.SQL.QueryRow(
					`SELECT contract_address, chain_id, abi from abis WHERE chain_id = $2 AND contract_address = $1`,
					"save_contract", "test",
				)
				err := row.Scan(&abi.ContractAddress, &abi.ChainID, &abi.ABI)

				suite.Require().NoError(err)
				suite.Require().Equal("save_contract", abi.ContractAddress)
				suite.Require().Equal("test", abi.ChainID)
				suite.Require().Equal([]byte(`{"name": "test"}`), abi.ABI)
			},
		},
		{
			name: "overwrites ABI correctly",
			setup: func(ctx context.Context) {
				err := suite.database.SaveAbi(ctx, evmtypes.NewAbi("test", "overwrite_contract", []byte(`{"name": "test"}`)))
				suite.Require().NoError(err)
			},
			abi:       evmtypes.NewAbi("test", "overwrite_contract", []byte(`{"name": "test"}`)),
			shouldErr: false,
			check: func() {
				var abi evmtypes.Abi
				row := suite.database.SQL.QueryRow(
					`SELECT contract_address, chain_id, abi from abis WHERE chain_id = $2 AND contract_address = $1`,
					"overwrite_contract", "test",
				)
				err := row.Scan(&abi.ContractAddress, &abi.ChainID, &abi.ABI)

				suite.Require().NoError(err)
				suite.Require().Equal("overwrite_contract", abi.ContractAddress)
				suite.Require().Equal("test", abi.ChainID)
				suite.Require().Equal([]byte(`{"name": "test"}`), abi.ABI)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			err := suite.database.SaveAbi(suite.T().Context(), tc.abi)
			if tc.shouldErr {
				suite.Assert().Error(err)
			} else {
				suite.Assert().NoError(err)
				if tc.check != nil {
					tc.check()
				}
			}
		})
	}
}

func (suite *DbTestSuite) TestGetABI() {
	testCases := []struct {
		name            string
		setup           func(ctx context.Context)
		contractAddress string
		chainID         string
		found           bool
		expected        *evmtypes.Abi
		shouldErr       bool
	}{
		{
			name: "Get ABI correctly",
			setup: func(ctx context.Context) {
				abi := evmtypes.NewAbi("polygon", "test", []byte("{}"))
				err := suite.database.SaveAbi(ctx, abi)
				suite.Require().NoError(err)
			},
			contractAddress: "test",
			chainID:         "polygon",
			found:           true,
			expected:        evmtypes.NewAbi("polygon", "test", []byte("{}")),
			shouldErr:       false,
		},
		{
			name:            "Not found ABI don't return error",
			contractAddress: "not_found",
			chainID:         "polygon",
			found:           false,
			shouldErr:       false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			abi, err := suite.database.GetAbi(suite.T().Context(), tc.chainID, tc.contractAddress)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			if tc.found {
				suite.Require().NotNil(abi)
				suite.Require().Equal(abi, tc.expected)
			} else {
				suite.Require().Nil(abi)
			}
		})
	}
}
