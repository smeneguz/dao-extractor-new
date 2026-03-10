package postgresql_test

import (
	"context"

	"github.com/dao-portal/extractor/database/postgresql"
	"github.com/dao-portal/extractor/types"
)

func (suite *DbTestSuite) TestInsertDAO() {
	testCases := []struct {
		name          string
		setup         func(ctx context.Context)
		dao           *types.DAO
		allowConflict bool
		shouldErr     bool
		check         func()
	}{
		{
			name:          "insert a non existing dao correctly",
			dao:           types.NewDAO("DAO", "Test DAO"),
			allowConflict: false,
			shouldErr:     false,
			check: func() {
				var dao postgresql.DAORow
				err := suite.database.SQL.Get(&dao,
					`SELECT * FROM daos WHERE symbol = $1`,
					"DAO",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.DAOID(1), dao.ID)
				suite.Require().Equal(types.DAOSymbol("DAO"), dao.Symbol)
				suite.Require().Equal("Test DAO", dao.Name)
			},
		},
		{
			name: "insert an existing dao fails",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertDAO(ctx, types.NewDAO("DAO", "Test DAO"), false)
				suite.Require().NoError(err)
			},
			dao:           types.NewDAO("DAO", "New DAO name"),
			allowConflict: false,
			shouldErr:     true,
		},
		{
			name: "insert an existing dao with allow conflict don not update the dao",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertDAO(ctx, types.NewDAO("DAO", "Test DAO"), false)
				suite.Require().NoError(err)
			},
			dao:           types.NewDAO("DAO", "New DAO"),
			allowConflict: true,
			shouldErr:     false,
			check: func() {
				var dao postgresql.DAORow
				err := suite.database.SQL.Get(&dao,
					`SELECT * FROM daos WHERE symbol = $1`,
					"DAO",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.DAOID(1), dao.ID)
				suite.Require().Equal(types.DAOSymbol("DAO"), dao.Symbol)
				suite.Require().Equal("Test DAO", dao.Name)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			_, err := suite.database.InsertDAO(suite.T().Context(), tc.dao, tc.allowConflict)
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

func (suite *DbTestSuite) TestUpdateDAO() {
	testCases := []struct {
		name      string
		setup     func(ctx context.Context)
		dao       *types.DAO
		shouldErr bool
		check     func()
	}{
		{
			name:      "update a non existing dao correctly fails",
			dao:       types.NewDAO("DAO", "Test DAO"),
			shouldErr: true,
		},
		{
			name: "update an existing dao correctly",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertDAO(ctx, types.NewDAO("DAO", "Test DAO"), false)
				suite.Require().NoError(err)
			},
			dao:       types.NewDAO("DAO", "New DAO name"),
			shouldErr: false,
			check: func() {
				var dao postgresql.DAORow
				err := suite.database.SQL.Get(&dao,
					`SELECT * FROM daos WHERE symbol = $1`,
					"DAO",
				)
				suite.Require().NoError(err)
				suite.Require().Equal(types.DAOID(1), dao.ID)
				suite.Require().Equal(types.DAOSymbol("DAO"), dao.Symbol)
				suite.Require().Equal("New DAO name", dao.Name)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			_, err := suite.database.UpdateDAO(suite.T().Context(), tc.dao)
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

func (suite *DbTestSuite) TestGetDAOBySymbol() {
	testCases := []struct {
		name      string
		setup     func(ctx context.Context)
		symbol    string
		shouldErr bool
		expected  *types.DAO
	}{
		{
			name:      "get a non existing dao correctly",
			symbol:    "DAO",
			shouldErr: false,
			expected:  nil,
		},
		{
			name: "get an existing dao correctly",
			setup: func(ctx context.Context) {
				_, err := suite.database.InsertDAO(ctx, types.NewDAO("DAO", "Test DAO"), false)
				suite.Require().NoError(err)
			},
			symbol:    "DAO",
			shouldErr: false,
			expected:  types.NewDAO("DAO", "Test DAO").WithID(1),
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			dao, err := suite.database.GetDAOBySymbol(suite.T().Context(), tc.symbol)
			if tc.shouldErr {
				suite.Require().Error(err)
				return
			}

			suite.Require().NoError(err)
			suite.Require().Equal(dao, tc.expected)
		})
	}
}
