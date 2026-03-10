package postgresql_test

import (
	"context"

	"github.com/dao-portal/extractor/database/postgresql"
	"github.com/dao-portal/extractor/types"
)

func (suite *DbTestSuite) TestInsertBlockchain() {
	testCases := []struct {
		name          string
		setup         func(ctx context.Context)
		blockchain    *types.Blockchain
		allowConflict bool
		shouldErr     bool
		check         func()
	}{
		{
			name:          "insert a non existing blockchain correctly",
			blockchain:    types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM),
			allowConflict: false,
			shouldErr:     false,
			check: func() {
				var blockchain postgresql.BlockChainRow
				err := suite.database.SQL.Get(&blockchain,
					`SELECT * FROM blockchains WHERE chain_id = $1`,
					"1",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.BlockchainID(1), blockchain.ID)
				suite.Require().Equal(types.ChainID("1"), blockchain.ChainID)
				suite.Require().Equal("Ethereum", blockchain.Name)
				suite.Require().Equal(types.ChainTypeEVM, blockchain.Type)
			},
		},
		{
			name: "insert an existing blockchain fails",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
			},
			blockchain:    types.NewBlockchain("1", "Ethereum NEW", types.ChainTypeEVM),
			allowConflict: false,
			shouldErr:     true,
		},
		{
			name: "insert an existing blockchain with allow conflict don not update the blockchain",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
			},
			blockchain:    types.NewBlockchain("1", "Ethereum NEW", types.ChainTypeEVM),
			allowConflict: true,
			shouldErr:     false,
			check: func() {
				var blockchain postgresql.BlockChainRow
				err := suite.database.SQL.Get(&blockchain,
					`SELECT * FROM blockchains WHERE chain_id = $1`,
					"1",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.BlockchainID(1), blockchain.ID)
				suite.Require().Equal(types.ChainID("1"), blockchain.ChainID)
				suite.Require().Equal("Ethereum", blockchain.Name)
				suite.Require().Equal(types.ChainTypeEVM, blockchain.Type)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			_, err := suite.database.InsertBlockchain(suite.T().Context(), tc.blockchain, tc.allowConflict)
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

func (suite *DbTestSuite) TestUpdateBlockchain() {
	testCases := []struct {
		name       string
		setup      func(ctx context.Context)
		blockchain *types.Blockchain
		shouldErr  bool
		check      func()
	}{
		{
			name:       "update a non existing blockchain correctly fails",
			blockchain: types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM),
			shouldErr:  true,
		},
		{
			name: "update an existing blockchain correctly",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
			},
			blockchain: types.NewBlockchain("1", "Ethereum NEW", types.ChainTypeEVM),
			shouldErr:  false,
			check: func() {
				var blockchain postgresql.BlockChainRow
				err := suite.database.SQL.Get(&blockchain,
					`SELECT * FROM blockchains WHERE chain_id = $1`,
					"1",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.BlockchainID(1), blockchain.ID)
				suite.Require().Equal(types.ChainID("1"), blockchain.ChainID)
				suite.Require().Equal("Ethereum NEW", blockchain.Name)
				suite.Require().Equal(types.ChainTypeEVM, blockchain.Type)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			_, err := suite.database.UpdateBlockchain(suite.T().Context(), tc.blockchain)
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

func (suite *DbTestSuite) TestGetBlockchainByChainID() {
	testCases := []struct {
		name      string
		setup     func(ctx context.Context)
		chainID   string
		shouldErr bool
		check     func()
		expected  *types.Blockchain
	}{
		{
			name:      "get a non existing blockchain correctly",
			chainID:   "1",
			shouldErr: false,
			expected:  nil,
		},
		{
			name: "get an existing blockchain correctly",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
			},
			chainID:   "1",
			shouldErr: false,
			expected:  types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM).WithID(1),
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			blockchain, err := suite.database.GetBlockchainByChainID(suite.T().Context(), tc.chainID)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			if tc.expected != nil {
				suite.Require().NotNil(blockchain)
				suite.Require().Equal(tc.expected, blockchain)
			} else {
				suite.Require().Nil(blockchain)
			}
		})
	}
}

func (suite *DbTestSuite) TestIsnertAddress() {
	testCases := []struct {
		name          string
		setup         func(ctx context.Context)
		address       *types.Address
		allowConflict bool
		shouldErr     bool
		check         func()
	}{
		{
			name:          "insert a non existing address correctly",
			address:       types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex),
			allowConflict: false,
			shouldErr:     false,
			check: func() {
				var address postgresql.AddressRow
				err := suite.database.SQL.Get(&address,
					`SELECT * FROM addresses WHERE address = $1`,
					"0x1",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.AddressID(1), address.ID)
				suite.Require().Equal("0x1", address.Address)
				suite.Require().Equal("label", address.Label)
				suite.Require().Equal(false, address.IsContract)
				suite.Require().Equal(types.AddressEncodingTypeHex, address.Encoding)
			},
		},
		{
			name: "insert an existing address fails",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertAddress(ctx, types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex), false)
				suite.Require().NoError(err)
			},
			address:       types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex),
			allowConflict: false,
			shouldErr:     true,
		},
		{
			name: "insert an existing address with allow conflict don not update the address",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertAddress(ctx, types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex), false)
				suite.Require().NoError(err)
			},
			address:       types.NewAddress("0x1", "label2", true, types.AddressEncodingTypeHex),
			allowConflict: true,
			shouldErr:     false,
			check: func() {
				var address postgresql.AddressRow
				err := suite.database.SQL.Get(&address,
					`SELECT * FROM addresses WHERE address = $1`,
					"0x1",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.AddressID(1), address.ID)
				suite.Require().Equal("0x1", address.Address)
				suite.Require().Equal("label", address.Label)
				suite.Require().Equal(false, address.IsContract)
				suite.Require().Equal(types.AddressEncodingTypeHex, address.Encoding)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			_, err := suite.database.InsertAddress(suite.T().Context(), tc.address, tc.allowConflict)
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

func (suite *DbTestSuite) TestUpdateAddress() {
	testCases := []struct {
		name      string
		setup     func(ctx context.Context)
		address   *types.Address
		shouldErr bool
		check     func()
	}{
		{
			name:      "update a non existing address correctly fails",
			address:   types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex),
			shouldErr: true,
		},
		{
			name: "update an existing address correctly",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertAddress(ctx, types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex), false)
				suite.Require().NoError(err)
			},
			address:   types.NewAddress("0x1", "label2", true, types.AddressEncodingTypeHex),
			shouldErr: false,
			check: func() {
				var address postgresql.AddressRow
				err := suite.database.SQL.Get(&address,
					`SELECT * FROM addresses WHERE address = $1`,
					"0x1",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.AddressID(1), address.ID)
				suite.Require().Equal(address.Address, "0x1")
				suite.Require().Equal(address.Label, "label2")
				suite.Require().Equal(address.IsContract, true)
				suite.Require().Equal(address.Encoding, types.AddressEncodingTypeHex)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			_, err := suite.database.UpdateAddress(suite.T().Context(), tc.address)
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

func (suite *DbTestSuite) TestGetAddressByAddress() {
	testCases := []struct {
		name      string
		setup     func(ctx context.Context)
		address   string
		shouldErr bool
		expected  *types.Address
	}{
		{
			name:      "get a non existing address correctly",
			address:   "0x1",
			shouldErr: false,
			expected:  nil,
		},
		{
			name: "get an existing address correctly",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertAddress(ctx, types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex), false)
				suite.Require().NoError(err)
			},
			address:   "0x1",
			shouldErr: false,
			expected:  types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex).WithID(1),
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			address, err := suite.database.GetAddressByAddress(suite.T().Context(), tc.address)
			if tc.shouldErr {
				suite.Require().Error(err)
				return
			}

			suite.Require().NoError(err)
			suite.Require().Equal(tc.expected, address)
		})
	}
}

func (suite *DbTestSuite) TestAssociateAddressToBlockchain() {
	testCases := []struct {
		name               string
		setup              func(ctx context.Context)
		address            *types.Address
		blockchain         *types.Blockchain
		shouldErr          bool
		check              func()
		expectedBlockchain *types.Blockchain
	}{
		{
			name: "associate an address to a non existing blockchain fails",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertAddress(ctx, types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex), false)
				suite.Require().NoError(err)
			},
			address:    types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex).WithID(1),
			blockchain: types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM).WithID(1),
			shouldErr:  true,
		},
		{
			name: "associating a non existing address to a blockchain fails",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
			},
			address:    types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex).WithID(1),
			blockchain: types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM).WithID(1),
			shouldErr:  true,
		},
		{
			name: "associate an address to a blockchain correctly",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertAddress(ctx, types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex), false)
				suite.Require().NoError(err)
				_, err = suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM), false)
				suite.Require().NoError(err)
			},
			address:    types.NewAddress("0x1", "label", false, types.AddressEncodingTypeHex).WithID(1),
			blockchain: types.NewBlockchain("1", "Ethereum", types.ChainTypeEVM).WithID(1),
			shouldErr:  false,
			check: func() {
				var addressBlockchains postgresql.AddressChainRow
				err := suite.database.SQL.Get(&addressBlockchains,
					`SELECT * FROM address_blockchains WHERE address_id = $1 AND blockchain_id = $2`,
					1, 1,
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.AddressID(1), addressBlockchains.AddressID)
				suite.Require().Equal(types.BlockchainID(1), addressBlockchains.BlockChainID)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			err := suite.database.AssociateAddressToBlockchain(suite.T().Context(), tc.address, tc.blockchain)
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
